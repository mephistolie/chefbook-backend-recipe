package dto

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func newProfilesInfo(profilesInfo map[string]entity.ProfileInfo) map[string]*api.RecipeProfileInfo {
	response := make(map[string]*api.RecipeProfileInfo)
	for id, profileInfo := range profilesInfo {
		response[id] = &api.RecipeProfileInfo{
			Name:   profileInfo.Name,
			Avatar: profileInfo.Avatar,
		}
	}
	return response
}
