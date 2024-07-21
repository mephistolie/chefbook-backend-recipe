package recipe

import (
	"context"
	"github.com/google/uuid"
	encryptionApi "github.com/mephistolie/chefbook-backend-encryption/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"sync"
	"time"
)

func (s *Service) GetRecipes(params entity.RecipesQuery, userId uuid.UUID, language string) entity.RecipesInfo {
	recipes := s.recipeRepo.GetRecipes(params, userId)

	var tagIds []string
	var collectionIds []uuid.UUID
	var profileIds []string
	for _, recipe := range recipes {
		tagIds = append(tagIds, recipe.Tags...)
		collectionIds = append(collectionIds, recipe.Collections...)
		profileIds = append(profileIds, recipe.OwnerId.String())
	}

	wg := sync.WaitGroup{}
	wg.Add(3)

	var tags map[string]entity.Tag
	var tagGroups map[string]string
	go func() {
		tags, tagGroups = s.getTags(tagIds, language)
		wg.Done()
	}()

	var collections map[uuid.UUID]entity.CollectionInfo
	go func() {
		collections = s.getCollectionsMap(collectionIds)
		wg.Done()
	}()

	var profilesInfo map[string]entity.ProfileInfo
	go func() {
		profilesInfo = s.getProfilesInfo(profileIds)
		wg.Done()
	}()

	wg.Wait()

	return entity.RecipesInfo{
		Recipes:      recipes,
		Collections:  collections,
		Tags:         tags,
		TagGroups:    tagGroups,
		ProfilesInfo: profilesInfo,
	}
}

func (s *Service) GetRecipesBook(userId uuid.UUID, language string) (entity.RecipeBook, error) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	collections := s.collectionRepo.GetCollections(userId, userId)

	var hasEncryptedVault bool
	go func() {
		ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
		res, err := s.grpc.Encryption.HasEncryptedVault(ctx, &encryptionApi.HasEncryptedVaultRequest{UserId: userId.String()})
		cancelCtx()
		if err == nil {
			hasEncryptedVault = res.HasEncryptedVault
		}
		wg.Done()
	}()

	recipes, err := s.recipeRepo.GetRecipeBook(userId)
	if err != nil {
		return entity.RecipeBook{}, err
	}

	var profiles []string
	var tagIds []string
	for _, recipe := range recipes {
		profiles = append(profiles, recipe.OwnerId.String())
		tagIds = append(tagIds, recipe.Tags...)
	}
	for _, collection := range collections {
		for _, contributor := range collection.Contributors {
			profiles = append(profiles, contributor.Id.String())
		}
	}

	var tags map[string]entity.Tag
	var tagGroups map[string]string
	go func() {
		tags, tagGroups = s.getTags(tagIds, language)
		wg.Done()
	}()

	profilesInfo := s.getProfilesInfo(profiles)

	wg.Wait()

	return entity.RecipeBook{
		Recipes:           recipes,
		Tags:              tags,
		TagGroups:         tagGroups,
		Collections:       collections,
		HasEncryptedVault: hasEncryptedVault,
		ProfilesInfo:      profilesInfo,
	}, nil
}

func (s *Service) GetRecipeNames(recipeIds []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]string, error) {
	return s.recipeRepo.GetRecipeNames(recipeIds, userId)
}
