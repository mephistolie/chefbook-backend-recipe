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

func (s *Service) SaveToRecipeBook(recipeId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.SaveToRecipeBook(recipeId, userId)
}

func (s *Service) RemoveFromRecipeBook(recipeId, userId uuid.UUID) error {
	return s.recipeRepo.RemoveFromRecipeBook(recipeId, userId)
}

func (s *Service) SetRecipeFavouriteStatus(recipeId, userId uuid.UUID, favourite bool) error {
	if err := s.checkRecipeAccessible(recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.SetRecipeFavouriteStatus(recipeId, userId, favourite)
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
