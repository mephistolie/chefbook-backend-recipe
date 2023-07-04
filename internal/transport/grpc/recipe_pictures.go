package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/grpc/dto"
)

func (s *RecipeServer) GenerateRecipePicturesUploadLinks(_ context.Context, req *api.GenerateRecipePicturesUploadLinksRequest) (*api.GenerateRecipePicturesUploadLinksResponse, error) {
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	picturesCount := int(req.PicturesCount)
	if picturesCount == 0 {
		return nil, fail.GrpcInvalidBody
	}

	maxPictures := s.subscriptionLimiter.GetMaxPicturesCount(req.Subscription)
	if picturesCount > maxPictures {
		picturesCount = maxPictures
	}

	uploads, err := s.service.GenerateRecipePicturesUploadLinks(recipeId, userId, picturesCount, req.Subscription)
	if err != nil {
		return nil, err
	}

	dtos := make([]*api.RecipePictureUploadLink, len(uploads))
	for i, upload := range uploads {
		dtos[i] = &api.RecipePictureUploadLink{
			Link:     upload.URL,
			FormData: upload.FormData,
		}
	}

	return &api.GenerateRecipePicturesUploadLinksResponse{Links: dtos}, nil
}

func (s *RecipeServer) SetRecipePictures(_ context.Context, req *api.SetRecipePicturesRequest) (*api.SetRecipePicturesResponse, error) {
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	pictures := dto.NewRecipePictures(req)

	version, err := s.service.SetRecipePictures(recipeId, userId, pictures, req.Version, req.Subscription)
	if err != nil {
		return nil, err
	}

	return &api.SetRecipePicturesResponse{Version: version}, nil
}
