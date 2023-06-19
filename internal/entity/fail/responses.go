package fail

import (
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

var (
	typeOutdatedVersion = "outdated_version"
)

var (
	GrpcOutdatedVersion = fail.CreateGrpcConflict(typeOutdatedVersion, "recipe version is outdated; process current version first")
)
