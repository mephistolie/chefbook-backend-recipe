package dto

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func newTags(tags map[string]entity.Tag) map[string]*api.RecipeTag {
	response := make(map[string]*api.RecipeTag)
	for id, tag := range tags {
		response[id] = newTag(tag)
	}
	return response
}

func newTag(tag entity.Tag) *api.RecipeTag {
	return &api.RecipeTag{
		Name:    tag.Name,
		Emoji:   tag.Emoji,
		GroupId: tag.GroupId,
	}
}
