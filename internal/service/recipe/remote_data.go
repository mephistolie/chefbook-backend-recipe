package recipe

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	slices "github.com/mephistolie/chefbook-backend-common/utils/slices"
	profileApi "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	tagApi "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	"time"
)

func (s *Service) getRecipeProfilesInfo(
	recipe entity.Recipe,
) map[string]entity.ProfileInfo {
	profiles := []string{recipe.OwnerId.String()}

	for i, _ := range recipe.Translations {
		for j, _ := range recipe.Translations[i] {
			profiles = append(profiles, recipe.Translations[i][j].String())
		}
	}

	return s.getProfilesInfo(profiles)
}

func (s *Service) getProfilesInfo(profileIds []string) map[string]entity.ProfileInfo {
	uniqueProfileIds := slices.RemoveDuplicates(profileIds)
	profilesInfo := make(map[string]entity.ProfileInfo)
	if len(uniqueProfileIds) == 0 {
		return profilesInfo
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := s.grpc.Profile.GetProfilesMinInfo(ctx, &profileApi.GetProfilesMinInfoRequest{ProfileIds: uniqueProfileIds})
	cancelCtx()

	if err == nil {
		for _, profileId := range uniqueProfileIds {
			if info, ok := res.Infos[profileId]; ok {
				profilesInfo[profileId] = entity.ProfileInfo{
					Name:   info.VisibleName,
					Avatar: info.Avatar,
				}
			}
		}
	} else {
		log.Warn("unable to get profiles info: %s", err)
	}

	return profilesInfo
}

func (s *Service) getTags(
	tagIds []string,
	languageCode string,
) (map[string]entity.Tag, map[string]string) {
	tags := make(map[string]entity.Tag)
	groups := make(map[string]string)

	uniqueTagIds := slices.RemoveDuplicates(tagIds)
	if len(uniqueTagIds) == 0 {
		return tags, groups
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := s.grpc.Tag.GetTagsMap(ctx, &tagApi.GetTagsMapRequest{
		TagIds:       uniqueTagIds,
		LanguageCode: languageCode,
	})
	cancelCtx()

	if err == nil {
		for tagId, dto := range res.Tags {
			tags[tagId] = entity.Tag{
				Id:      tagId,
				Name:    dto.Name,
				Emoji:   dto.Emoji,
				GroupId: dto.GroupId,
			}
		}
		for groupId, groupName := range res.GroupNames {
			groups[groupId] = groupName
		}
	} else {
		log.Warn("unable to get recipe tags: ", err)
	}

	return tags, groups
}

func (s *Service) getCollectionsMap(
	collectionIds []uuid.UUID,
) map[uuid.UUID]entity.CollectionInfo {
	collections := make(map[uuid.UUID]entity.CollectionInfo)

	collectionIdsMap := make(map[uuid.UUID]bool)
	var uniqueCollectionIds []uuid.UUID
	for _, collectionId := range collectionIds {
		if _, ok := collectionIdsMap[collectionId]; !ok {
			uniqueCollectionIds = append(uniqueCollectionIds, collectionId)
			collectionIdsMap[collectionId] = true
		}
	}

	if len(uniqueCollectionIds) > 0 {
		collections = s.collectionRepo.GetCollectionsMap(collectionIds)
	}

	return collections
}
