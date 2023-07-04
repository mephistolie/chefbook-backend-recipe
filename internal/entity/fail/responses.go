package fail

import (
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

var (
	typeOutdatedVersion = "outdated_version"
)

var (
	GrpcEmptyIngredients        = fail.CreateGrpcClient(fail.TypeInvalidBody, "ingredients mustn't be empty")
	GrpcEmptyCooking            = fail.CreateGrpcClient(fail.TypeInvalidBody, "cooking mustn't be empty")
	GrpcIngredientMatchingIds   = fail.CreateGrpcClient(fail.TypeInvalidBody, "ingredients have same ID")
	GrpcCookingMatchingIds      = fail.CreateGrpcClient(fail.TypeInvalidBody, "cooking items have same ID")
	GrpcEmptyIngredientText     = fail.CreateGrpcClient(fail.TypeInvalidBody, "ingredient text can't be empty")
	GrpcEmptyCookingItemText    = fail.CreateGrpcClient(fail.TypeInvalidBody, "cooking item text can't be empty")
	GrpcInvalidIngredientType   = fail.CreateGrpcClient(fail.TypeInvalidBody, "invalid ingredient type")
	GrpcInvalidCookingItemType  = fail.CreateGrpcClient(fail.TypeInvalidBody, "invalid cooking type")
	GrpcInvalidIngredientId     = fail.CreateGrpcClient(fail.TypeInvalidBody, "invalid ingredient ID")
	GrpcInvalidCookingItemId    = fail.CreateGrpcClient(fail.TypeInvalidBody, "invalid cooking ID")
	GrpcEncryptedPublicRecipe   = fail.CreateGrpcClient(fail.TypeInvalidBody, "encrypted recipe can't be public")
	GrpcInvalidEncryptedFormat  = fail.CreateGrpcClient(fail.TypeInvalidBody, "invalid format for encrypted recipe")
	GrpcChangedEncryptionStatus = fail.CreateGrpcClient(fail.TypeInvalidBody, "changing encryption status is forbidden")

	GrpcRecipeExists    = fail.CreateGrpcConflict(typeOutdatedVersion, "recipe with specified ID already exists")
	GrpcOutdatedVersion = fail.CreateGrpcConflict(typeOutdatedVersion, "recipe version is outdated; process current version first")

	GrpcRecipePictureNotFound    = fail.CreateGrpcNotFound(fail.TypeNotFound, "one or more entered pictures hasn't been found")
	GrpcRecipePicturesCountLimit = fail.CreateGrpcAccessDenied(fail.TypeAccessDenied, "recipe pictures count out of limit")
)
