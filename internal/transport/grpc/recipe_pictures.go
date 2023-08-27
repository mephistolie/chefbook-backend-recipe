package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
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
	if picturesCount <= 0 {
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
			PictureLink: upload.PictureLink,
			UploadLink:  upload.UploadUrl,
			FormData:    upload.FormData,
			MaxSize:     upload.MaxSize,
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
	if !s.checkForPictureDuplicates(pictures) {
		return nil, recipeFail.GrpcDuplicatePictures
	}

	version, err := s.service.SetRecipePictures(recipeId, userId, pictures, req.Version, req.Subscription)
	if err != nil {
		return nil, err
	}

	return &api.SetRecipePicturesResponse{Version: version}, nil
}

func (s *RecipeServer) checkForPictureDuplicates(pictures entity.RecipePictures) bool {
	usedIds := make(map[string]bool)
	if pictures.Preview != nil {
		usedIds[*pictures.Preview] = true
	}
	for _, stepPictures := range pictures.Cooking {
		for _, stepPicture := range stepPictures {
			if exists, ok := usedIds[stepPicture]; ok && exists {
				return false
			}
			usedIds[stepPicture] = true
		}
	}
	return true
}
