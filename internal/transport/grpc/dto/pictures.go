package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func NewRecipePictures(req *api.SetRecipePicturesRequest) entity.RecipePictures {
	pictures := entity.RecipePictures{}

	if req.Preview != nil {
		if previewId, err := uuid.Parse(*req.Preview); err == nil {
			pictures.Preview = &previewId
		}
	}
	if len(req.CookingPicturesMap) > 0 {
		cookingPicturesMap := make(map[uuid.UUID][]uuid.UUID)
		for rawStepId, rawStepPictures := range req.CookingPicturesMap {
			if stepId, err := uuid.Parse(rawStepId); err == nil {
				var stepPictures []uuid.UUID
				for _, rawPictureId := range rawStepPictures.Pictures {
					if pictureId, err := uuid.Parse(rawPictureId); err == nil {
						stepPictures = append(stepPictures, pictureId)
					}
				}
				cookingPicturesMap[stepId] = stepPictures
			}
		}
		pictures.Cooking = cookingPicturesMap
	}

	return pictures
}
