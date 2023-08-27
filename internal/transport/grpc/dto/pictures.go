package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func NewRecipePictures(req *api.SetRecipePicturesRequest) entity.RecipePictures {
	pictures := entity.RecipePictures{Preview: req.Preview}

	if len(req.CookingPicturesMap) > 0 {
		cookingPicturesMap := make(map[uuid.UUID][]string)
		for rawStepId, stepPictures := range req.CookingPicturesMap {
			if stepId, err := uuid.Parse(rawStepId); err == nil {
				cookingPicturesMap[stepId] = stepPictures.Pictures
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
