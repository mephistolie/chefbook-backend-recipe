package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func NewRecipeTranslations(translations map[string][]entity.RecipeTranslationInfo) map[string]*api.RecipeTranslations {
	dto := map[string]*api.RecipeTranslations{}
	for language, languageTranslations := range translations {
		var dtos []*api.RecipeTranslationInfo
		for _, translation := range languageTranslations {
			dtos = append(dtos, &api.RecipeTranslationInfo{
				AuthorId:     translation.AuthorId.String(),
				AuthorName:   translation.AuthorName,
				AuthorAvatar: translation.AuthorAvatar,
			})
		}
		dto[language] = &api.RecipeTranslations{Translations: dtos}
	}
	return dto
}

func NewRecipeTranslation(req *api.TranslateRecipeRequest) (entity.RecipeTranslation, error) {
	translatorId, err := uuid.Parse(req.TranslatorId)
	if err != nil {
		return entity.RecipeTranslation{}, fail.GrpcInvalidBody
	}
	if len(req.Language) != 2 || len(req.Name) == 0 || len(req.Ingredients) == 0 || len(req.Cooking) == 0 {
		return entity.RecipeTranslation{}, fail.GrpcInvalidBody
	}
	if len(req.Name) > maxNameLength {
		req.Name = req.Name[0:maxNameLength]
	}
	if req.Description != nil && len(*req.Description) > maxDescriptionLength {
		description := (*req.Description)[0:maxDescriptionLength]
		req.Description = &description
	}

	ingredients, err := NewRecipeIngredientsTranslation(req)
	if err != nil {
		return entity.RecipeTranslation{}, err
	}
	cooking, err := NewRecipeCookingTranslation(req)
	if err != nil {
		return entity.RecipeTranslation{}, err
	}

	return entity.RecipeTranslation{
		AuthorId:    translatorId,
		Language:    req.Language,
		Name:        req.Name,
		Description: req.Description,
		Ingredients: ingredients,
		Cooking:     cooking,
	}, nil
}

func NewRecipeIngredientsTranslation(req *api.TranslateRecipeRequest) (map[uuid.UUID]entity.IngredientTranslation, error) {
	if len(req.Ingredients) > maxIngredientsCount {
		return nil, fail.GrpcInvalidBody
	}

	ingredients := map[uuid.UUID]entity.IngredientTranslation{}
	for rawId, translation := range req.Ingredients {
		id, err := uuid.Parse(rawId)
		if err != nil {
			return nil, fail.GrpcInvalidBody
		}
		ingredients[id] = entity.IngredientTranslation{
			Text: translation.Text,
			Unit: translation.Unit,
		}
	}
	return ingredients, nil
}

func NewRecipeCookingTranslation(req *api.TranslateRecipeRequest) (map[uuid.UUID]string, error) {
	if len(req.Cooking) > maxCookingStepsCount {
		return nil, fail.GrpcInvalidBody
	}

	cooking := map[uuid.UUID]string{}
	for rawId, translation := range req.Cooking {
		id, err := uuid.Parse(rawId)
		if err != nil {
			return nil, fail.GrpcInvalidBody
		}
		cooking[id] = translation
	}
	return cooking, nil
}
