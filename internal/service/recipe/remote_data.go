package recipe

import (
	"context"
	"github.com/google/uuid"
	categoryApi "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/log"
	slices "github.com/mephistolie/chefbook-backend-common/utils/slices"
	profileApi "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	tagApi "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	"sync"
	"time"
)

func (s *Service) getRecipeAuthorsInfo(authorIds []string) map[string]*profileApi.ProfileMinInfo {
	uniqueAuthorIds := slices.RemoveDuplicates(authorIds)
	infos := make(map[string]*profileApi.ProfileMinInfo)
	if len(uniqueAuthorIds) == 0 {
		return infos
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := s.grpc.Profile.GetProfilesMinInfo(ctx, &profileApi.GetProfilesMinInfoRequest{ProfileIds: uniqueAuthorIds})
	cancelCtx()

	if err == nil {
		for _, authorId := range uniqueAuthorIds {
			if info, ok := res.Infos[authorId]; ok {
				infos[authorId] = info
			}
		}
	} else {
		log.Warn("unable to get recipe authors data: %s", err)
	}

	return infos
}

func (s *Service) getCategoriesAndTagsAsync(
	tagIds []string,
	categoryIds []uuid.UUID,
	userId uuid.UUID,
	language string,
	tagsDestination *map[string]entity.Tag,
	categoriesDestination *map[string]entity.Category,
) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var rawCategoryIds []string
	for _, categoryId := range categoryIds {
		rawCategoryIds = append(rawCategoryIds, categoryId.String())
	}

	go s.getTags(language, tagIds, tagsDestination, &wg)
	go s.getCategoriesMap(userId, rawCategoryIds, categoriesDestination, &wg)

	return &wg
}

func (s *Service) getTags(
	languageCode string,
	tagIds []string,
	destination *map[string]entity.Tag,
	wg *sync.WaitGroup,
) {
	uniqueTagIds := slices.RemoveDuplicates(tagIds)
	if len(uniqueTagIds) == 0 {
		wg.Done()
		return
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := s.grpc.Tag.GetTagsMap(ctx, &tagApi.GetTagsMapRequest{
		TagIds:       uniqueTagIds,
		LanguageCode: languageCode,
	})
	cancelCtx()

	if err == nil {
		log.Debugf("found %d tags", len(res.Tags))
		for tagId, dto := range res.Tags {
			(*destination)[tagId] = entity.Tag{
				Id:    tagId,
				Name:  dto.Name,
				Emoji: dto.Emoji,
			}
		}
	} else {
		log.Warn("unable to get recipe tags: ", err)
	}

	wg.Done()
}

func (s *Service) getUserCategories(
	userId uuid.UUID,
) []entity.Category {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := s.grpc.Category.GetUserCategories(ctx, &categoryApi.GetUserCategoriesRequest{UserId: userId.String()})
	cancelCtx()

	var categories []entity.Category

	if err == nil {
		for _, dto := range res.Categories {
			categories = append(categories, entity.Category{
				Id:    dto.CategoryId,
				Name:  dto.Name,
				Emoji: dto.Emoji,
			})
		}
	}

	return categories
}

func (s *Service) getCategoriesMap(
	userId uuid.UUID,
	categoryIds []string,
	destination *map[string]entity.Category,
	wg *sync.WaitGroup,
) {
	uniqueCategoryIds := slices.RemoveDuplicates(categoryIds)
	if len(uniqueCategoryIds) == 0 {
		wg.Done()
		return
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := s.grpc.Category.GetCategoriesMap(ctx, &categoryApi.GetCategoriesMapRequest{
		CategoryIds: uniqueCategoryIds,
		UserId:      userId.String(),
	})
	cancelCtx()

	if err == nil {
		for categoryId, dto := range res.Categories {
			(*destination)[categoryId] = entity.Category{
				Id:    categoryId,
				Name:  dto.Name,
				Emoji: dto.Emoji,
			}
		}
	}

	wg.Done()
}
