package collection

import (
	"context"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/utils/slices"
	profileApi "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"time"
)

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
