package recipe

import (
	"context"
	"github.com/google/uuid"
	encryptionApi "github.com/mephistolie/chefbook-backend-encryption/api/proto/implementation/v1"
	profileApi "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"sync"
	"time"
)

func (s *Service) GetRecipes(params entity.RecipesQuery, userId uuid.UUID, language string) entity.DetailedRecipesInfo {
	recipes := s.repo.GetRecipes(params, userId)

	var authors []string
	var tagIds []string
	var categoryIds []uuid.UUID
	for _, recipe := range recipes {
		authors = append(authors, recipe.OwnerId.String())
		tagIds = append(tagIds, recipe.Tags...)
		categoryIds = append(categoryIds, recipe.Categories...)
	}

	tags := make(map[string]entity.Tag)
	tagGroups := make(map[string]string)
	categories := make(map[string]entity.Category)
	wg := s.getCategoriesAndTagsAsync(tagIds, categoryIds, userId, language, &tags, &tagGroups, &categories)

	authorsMap := s.getProfilesInfo(authors)

	wg.Wait()

	return s.getRecipeInfos(recipes, authorsMap, tags, tagGroups, categories)
}

func (s *Service) getRecipeInfos(
	recipes []entity.BaseRecipeInfo,
	authors map[string]*profileApi.ProfileMinInfo,
	tags map[string]entity.Tag,
	tagGroups map[string]string,
	categories map[string]entity.Category,
) entity.DetailedRecipesInfo {
	var infos []entity.RecipeInfo
	for _, baseRecipe := range recipes {
		recipe := entity.RecipeInfo{BaseRecipeInfo: baseRecipe}

		if info, ok := authors[recipe.OwnerId.String()]; ok && info != nil {
			recipe.OwnerName = info.VisibleName
			recipe.OwnerAvatar = info.Avatar
		}

		infos = append(infos, recipe)
	}

	return entity.DetailedRecipesInfo{
		Recipes:    infos,
		Categories: categories,
		Tags:       tags,
		TagGroups:  tagGroups,
	}
}

func (s *Service) GetRecipesBook(userId uuid.UUID, language string) (entity.DetailedRecipesState, error) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	var categories []entity.Category
	go func() {
		categories = s.getUserCategories(userId)
		wg.Done()
	}()

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

	recipes, err := s.repo.GetRecipeBook(userId)
	if err != nil {
		return entity.DetailedRecipesState{}, err
	}

	var authors []string
	var tagIds []string
	var categoryIds []uuid.UUID
	for _, recipe := range recipes {
		authors = append(authors, recipe.OwnerId.String())
		tagIds = append(tagIds, recipe.Tags...)
		categoryIds = append(categoryIds, recipe.Categories...)
	}

	tags := make(map[string]entity.Tag)
	tagGroups := make(map[string]string)
	go s.getTags(language, tagIds, &tags, &tagGroups, &wg)

	authorsMap := s.getProfilesInfo(authors)

	wg.Wait()

	return s.getRecipeStates(recipes, authorsMap, tags, tagGroups, categories, hasEncryptedVault), nil
}

func (s *Service) getRecipeStates(
	recipes []entity.BaseRecipeState,
	authors map[string]*profileApi.ProfileMinInfo,
	tags map[string]entity.Tag,
	tagGroups map[string]string,
	categories []entity.Category,
	hasEncryptedVault bool,
) entity.DetailedRecipesState {
	var states []entity.RecipeState
	for _, baseRecipe := range recipes {
		recipe := entity.RecipeState{BaseRecipeState: baseRecipe}

		if info, ok := authors[recipe.OwnerId.String()]; ok && info != nil {
			recipe.OwnerName = info.VisibleName
			recipe.OwnerAvatar = info.Avatar
		}

		states = append(states, recipe)
	}

	return entity.DetailedRecipesState{
		Recipes:           states,
		Tags:              tags,
		TagGroups:         tagGroups,
		Categories:        categories,
		HasEncryptedVault: hasEncryptedVault,
	}
}

func (s *Service) GetRecipeNames(recipeIds []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]string, error) {
	return s.repo.GetRecipeNames(recipeIds, userId)
}
