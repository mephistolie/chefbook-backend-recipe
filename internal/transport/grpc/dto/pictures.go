package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func NewRecipePictures(req *api.SetRecipePicturesRequest) entity.RecipePictureIds {
	pictures := entity.RecipePictureIds{}

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

func newRecipesResponse(res entity.RecipePictures) *api.RecipePictures {
	var cooking map[string]*api.StepPictures
	for stepId, pictures := range res.Cooking {
		cooking[stepId.String()] = &api.StepPictures{Pictures: pictures}
	}

	pictures := api.RecipePictures{
		Preview: res.Preview,
		Cooking: cooking,
	}
	return &pictures
}
