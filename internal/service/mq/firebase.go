package mq

import (
	"context"
	"errors"
	"github.com/google/uuid"
	categoryApi "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/firebase"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"regexp"
	"strconv"
)

func (s *Service) ImportFirebaseRecipes(userId uuid.UUID, firebaseId string, messageId uuid.UUID) error {
	if s.firebase == nil {
		log.Warnf("try to import firebase profile with firebase import disabled")
		return errors.New("firebase import disabled")
	}

	firebaseRecipes, err := s.firebase.GetRecipes(firebaseId)
	if err != nil {
		log.Warnf("unable to get firebase recipes for user %s: %s", userId, err)
		return err
	}

	log.Infof("found %d Firebase recipes for user %s...", len(firebaseRecipes), userId)
	if err := s.repo.ConfirmFirebaseDataLoad(messageId); err != nil {
		return err
	}

	categories := make(map[string]uuid.UUID)
	for _, firebaseRecipe := range firebaseRecipes {
		s.importFirebaseRecipe(firebaseRecipe, userId, &categories)
	}

	return nil
}

func (s *Service) importFirebaseRecipe(firebaseRecipe firebase.Recipe, userId uuid.UUID, categories *map[string]uuid.UUID) {
	var servingsPtr *int32
	if firebaseRecipe.Servings != nil {
		servings := int32(*firebaseRecipe.Servings)
		servingsPtr = &servings
	}
	var caloriesPtr *int32
	if firebaseRecipe.Calories != nil {
		calories := int32(*firebaseRecipe.Calories)
		caloriesPtr = &calories
	}

	recipe := entity.RecipeInput{
		UserId:            userId,
		Name:              firebaseRecipe.Name,
		Visibility:        model.VisibilityPrivate,
		IsEncrypted:       false,
		Language:          "en",
		Servings:          servingsPtr,
		Time:              parseRecipeTime(firebaseRecipe.Time),
		Calories:          caloriesPtr,
		Ingredients:       parseRecipeIngredients(firebaseRecipe.Ingredients),
		Cooking:           parseRecipeCooking(firebaseRecipe.Cooking),
		CreationTimestamp: firebaseRecipe.CreationTimestamp,
	}

	var recipeCategories []uuid.UUID
	for _, category := range firebaseRecipe.Categories {
		if categoryId, ok := (*categories)[category]; ok {
			recipeCategories = append(recipeCategories, categoryId)
		} else {
			if res, err := s.grpc.Category.CreateCategory(context.Background(), &categoryApi.CreateCategoryRequest{
				Name:   category,
				UserId: userId.String(),
			}); err == nil {
				if categoryId, err = uuid.Parse(res.CategoryId); err == nil {
					(*categories)[category] = categoryId
					recipeCategories = append(recipeCategories, categoryId)
				}
			}
		}
	}

	if recipeId, _, err := s.repo.CreateRecipe(recipe); err == nil {
		if firebaseRecipe.IsFavourite {
			_ = s.repo.SetRecipeFavouriteStatus(recipeId, userId, true)
		}
		if len(recipeCategories) > 0 {
			_ = s.repo.SetRecipeCategories(recipeId, userId, recipeCategories)
		}
	}
}

func parseRecipeTime(timeString *string) *int32 {
	if timeString == nil {
		return nil
	}

	minutes := 0
	numberFilter := regexp.MustCompile("[0-9]+")
	timeSlice := numberFilter.FindAllString(*timeString, -1)
	timeSliceLength := len(timeSlice)
	if timeSliceLength > 0 {
		multiplier := 1
		if timeSliceLength == 1 && len(timeSlice[timeSliceLength-1]) == 1 {
			multiplier = 60
		}
		number, err := strconv.Atoi(timeSlice[timeSliceLength-1])
		if err == nil {
			minutes += number * multiplier
		}
	}
	if timeSliceLength > 1 {
		hours, err := strconv.Atoi(timeSlice[timeSliceLength-2])
		if err == nil {
			minutes += hours * 60
		}
	}
	if minutes > 0 {
		minutes32 := int32(minutes)
		return &minutes32
	}
	return nil
}

func parseRecipeIngredients(firebaseIngredients []firebase.Ingredient) []entity.IngredientItem {
	ingredients := make([]entity.IngredientItem, len(firebaseIngredients))
	for i, firebaseIngredient := range firebaseIngredients {
		text := firebaseIngredient.Text
		ingredient := entity.IngredientItem{
			Id:   uuid.New(),
			Text: &text,
			Type: entity.TypeIngredient,
		}
		if firebaseIngredient.Section {
			ingredient.Type = entity.TypeIngredientsSection
		}

		ingredients[i] = ingredient
	}

	return ingredients
}

func parseRecipeCooking(firebaseCooking []firebase.Step) []entity.CookingItem {
	cooking := make([]entity.CookingItem, len(firebaseCooking))
	for i, firebaseCookingItem := range firebaseCooking {
		text := firebaseCookingItem.Text
		cookingItem := entity.CookingItem{
			Id:   uuid.New(),
			Text: &text,
			Type: entity.TypeStep,
		}
		if firebaseCookingItem.Section {
			cookingItem.Type = entity.TypeCookingSection
		}

		cooking[i] = cookingItem
	}

	return cooking
}
