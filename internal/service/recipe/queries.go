package recipe

import (
	"github.com/google/uuid"
	profileApi "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"sync"
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
	categories := make(map[string]entity.Category)
	wg := s.getCategoriesAndTagsAsync(tagIds, categoryIds, userId, language, &tags, &categories)

	authorsMap := s.getRecipeAuthorsInfo(authors)

	wg.Wait()

	return s.getRecipeInfos(recipes, authorsMap, tags, categories)
}

func (s *Service) getRecipeInfos(
	recipes []entity.BaseRecipeInfo,
	authors map[string]*profileApi.ProfileMinInfo,
	tags map[string]entity.Tag,
	categories map[string]entity.Category,
) entity.DetailedRecipesInfo {
	var infos []entity.RecipeInfo
	for _, baseRecipe := range recipes {
		recipe := entity.RecipeInfo{
			BaseRecipeInfo: baseRecipe,
		}

		if info, ok := authors[recipe.OwnerId.String()]; ok && info != nil {
			if len(info.VisibleName) > 0 {
				recipe.OwnerName = &info.VisibleName
			}
			if len(info.Avatar) > 0 {
				recipe.OwnerAvatar = &info.Avatar
			}
		}

		infos = append(infos, recipe)
	}

	return entity.DetailedRecipesInfo{
		Recipes:    infos,
		Tags:       tags,
		Categories: categories,
	}
}

func (s *Service) GetRecipesBook(userId uuid.UUID, language string) (entity.DetailedRecipesState, error) {
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
	var categories []entity.Category

	wg := sync.WaitGroup{}
	wg.Add(2)

	go s.getTags(language, tagIds, &tags, &wg)
	go func() {
		categories = s.getUserCategories(userId)
		wg.Done()
	}()

	authorsMap := s.getRecipeAuthorsInfo(authors)

	wg.Wait()

	return s.getRecipeStates(recipes, authorsMap, tags, categories), nil
}

func (s *Service) getRecipeStates(
	recipes []entity.BaseRecipeState,
	authors map[string]*profileApi.ProfileMinInfo,
	tags map[string]entity.Tag,
	categories []entity.Category,
) entity.DetailedRecipesState {
	var states []entity.RecipeState
	for _, baseRecipe := range recipes {
		recipe := entity.RecipeState{
			BaseRecipeState: baseRecipe,
		}

		if info, ok := authors[recipe.OwnerId.String()]; ok && info != nil {
			if len(info.VisibleName) > 0 {
				recipe.OwnerName = &info.VisibleName
			}
			if len(info.Avatar) > 0 {
				recipe.OwnerAvatar = &info.Avatar
			}
		}

		states = append(states, recipe)
	}

	return entity.DetailedRecipesState{
		Recipes:    states,
		Tags:       tags,
		Categories: categories,
	}
}
