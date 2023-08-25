package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
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
			PictureId:   upload.PictureId.String(),
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

	usedIds := make(map[uuid.UUID]bool)
	for _, id := range pictures.GetIds() {
		if exists, ok := usedIds[id]; ok && exists {
			return nil, recipeFail.GrpcDuplicatePictures
		}
		usedIds[id] = true
	}

	links, version, err := s.service.SetRecipePictures(recipeId, userId, pictures, req.Version, req.Subscription)
	if err != nil {
		return nil, err
	}

	rawLinks := make(map[string]string)
	for pictureId, link := range links {
		rawLinks[pictureId.String()] = link
	}

	return &api.SetRecipePicturesResponse{Links: rawLinks, Version: version}, nil
}
