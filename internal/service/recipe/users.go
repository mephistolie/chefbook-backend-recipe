package recipe

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
)

func (s *Service) RateRecipe(recipeId, userId uuid.UUID, score int) error {
	policy, err := s.recipeRepo.GetRecipePolicy(recipeId)
	if err != nil {
		return err
	}
	if policy.OwnerId == userId || policy.Visibility != model.VisibilityPublic {
		return fail.GrpcAccessDenied
	}

	msg, err := s.recipeRepo.RateRecipe(recipeId, userId, score)
	if err == nil && msg != nil {
		go s.mqPublisher.PublishMessage(msg)
	}
	return err
}

func (s *Service) SaveRecipeToRecipeBook(recipeId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.SaveRecipeToRecipeBook(recipeId, userId)
}

func (s *Service) RemoveRecipeFromRecipeBook(recipeId, userId uuid.UUID) error {
	return s.recipeRepo.RemoveRecipeFromRecipeBook(recipeId, userId)
}

func (s *Service) SaveRecipeToFavourites(recipeId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.SaveRecipeToFavourites(recipeId, userId)
}

func (s *Service) RemoveRecipeFromFavourites(recipeId, userId uuid.UUID) error {
	return s.recipeRepo.RemoveRecipeFromFavourites(recipeId, userId)
}

func (s *Service) AddRecipeToCollection(recipeId, collectionId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.AddRecipeToCollection(recipeId, collectionId, userId)
}

func (s *Service) RemoveRecipeFromCollection(recipeId, collectionId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.AddRecipeToCollection(recipeId, collectionId, userId)
}

func (s *Service) SetRecipeCollections(recipeId, userId uuid.UUID, collections []uuid.UUID) error {
	if err := s.checkRecipeAccessible(recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.SetRecipeCollections(recipeId, userId, collections)
}

func (s *Service) checkRecipeAccessible(recipeId, userId uuid.UUID) error {
	policy, err := s.recipeRepo.GetRecipePolicy(recipeId)
	if err != nil {
		return err
	}

	if userId != policy.OwnerId && policy.Visibility == model.VisibilityPrivate {
		return fail.GrpcAccessDenied
	}

	return nil
}
