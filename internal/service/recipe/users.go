package recipe

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
)

func (s *Service) RateRecipe(ctx context.Context, recipeId, userId uuid.UUID, score int) error {
	policy, err := s.recipeRepo.GetRecipePolicy(ctx, recipeId)
	if err != nil {
		return err
	}
	if policy.OwnerId == userId || policy.Visibility != model.VisibilityPublic {
		return fail.GrpcAccessDenied
	}

	msg, err := s.recipeRepo.RateRecipe(ctx, recipeId, userId, score)
	if err == nil && msg != nil {
		go s.mqPublisher.PublishMessage(msg)
	}
	return err
}

func (s *Service) SaveRecipeToRecipeBook(ctx context.Context, recipeId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(ctx, recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.SaveRecipeToRecipeBook(ctx, recipeId, userId)
}

func (s *Service) RemoveRecipeFromRecipeBook(ctx context.Context, recipeId, userId uuid.UUID) error {
	return s.recipeRepo.RemoveRecipeFromRecipeBook(ctx, recipeId, userId)
}

func (s *Service) SaveRecipeToFavourites(ctx context.Context, recipeId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(ctx, recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.SaveRecipeToFavourites(ctx, recipeId, userId)
}

func (s *Service) RemoveRecipeFromFavourites(ctx context.Context, recipeId, userId uuid.UUID) error {
	return s.recipeRepo.RemoveRecipeFromFavourites(ctx, recipeId, userId)
}

func (s *Service) AddRecipeToCollection(ctx context.Context, recipeId, collectionId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(ctx, recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.AddRecipeToCollection(ctx, recipeId, collectionId, userId)
}

func (s *Service) RemoveRecipeFromCollection(ctx context.Context, recipeId, collectionId, userId uuid.UUID) error {
	if err := s.checkRecipeAccessible(ctx, recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.RemoveRecipeFromCollection(ctx, recipeId, collectionId, userId)
}

func (s *Service) SetRecipeCollections(ctx context.Context, recipeId, userId uuid.UUID, collections []uuid.UUID) error {
	if err := s.checkRecipeAccessible(ctx, recipeId, userId); err != nil {
		return err
	}
	return s.recipeRepo.SetRecipeCollections(ctx, recipeId, userId, collections)
}

func (s *Service) checkRecipeAccessible(ctx context.Context, recipeId, userId uuid.UUID) error {
	policy, err := s.recipeRepo.GetRecipePolicy(ctx, recipeId)
	if err != nil {
		return err
	}

	if userId != policy.OwnerId && policy.Visibility == model.VisibilityPrivate {
		return fail.GrpcAccessDenied
	}

	return nil
}
