package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres/dto"
)

func (r *Repository) GetRecipePictureIdsToUpload(recipeId uuid.UUID, picturesCount int) ([]uuid.UUID, error) {
	tx, err := r.startTransaction()
	if err != nil {
		return nil, err
	}

	pictures, hasRequest, err := r.getExistingRecipePicturesUploadsQuery(recipeId, tx)
	if err != nil {
		return nil, err
	}
	if !hasRequest {
		return r.createRecipePicturesUploadRequest(recipeId, picturesCount, tx)
	}

	existingUploadsSize := len(pictures)
	if existingUploadsSize > picturesCount {
		pictures = pictures[0:picturesCount]
	}

	if existingUploadsSize < picturesCount {
		for i := existingUploadsSize; i < picturesCount; i++ {
			pictures = append(pictures, uuid.New())
		}
		if err = r.updateRecipePicturesUploadRequest(recipeId, pictures, tx); err != nil {
			return nil, err
		}
	}

	return pictures, commitTransaction(tx)
}

func (r *Repository) getExistingRecipePicturesUploadsQuery(recipeId uuid.UUID, tx *sql.Tx) (dto.RecipePicturesUpload, bool, error) {
	var pictures dto.RecipePicturesUpload
	hasRequest := false

	getExistingPicturesQuery := fmt.Sprintf(`
		SELECT pictures
		FROM %[1]v
		WHERE recipe_id=$1
	`, recipePicturesUploadsTable)

	rows, err := tx.Query(getExistingPicturesQuery, recipeId)
	if err != nil {
		return nil, false, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	for rows.Next() {
		if err = rows.Scan(&pictures); err == nil {
			_ = rows.Close()
			hasRequest = true
			break
		}
	}

	return pictures, hasRequest, nil
}

func (r *Repository) createRecipePicturesUploadRequest(recipeId uuid.UUID, picturesCount int, tx *sql.Tx) ([]uuid.UUID, error) {
	var pictures dto.RecipePicturesUpload

	for i := 0; i < picturesCount; i++ {
		pictures = append(pictures, uuid.New())
	}

	createRequestQuery := fmt.Sprintf(`
		INSERT INTO %[1]v (recipe_id, pictures)
		VALUES ($1, $2)
	`, recipePicturesUploadsTable)

	if _, err := tx.Exec(createRequestQuery, recipeId, pictures); err != nil {
		log.Errorf("unable to create recipe %s pictures uploading request: %s", recipeId, err)
		return nil, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}
	return pictures, commitTransaction(tx)
}

func (r *Repository) updateRecipePicturesUploadRequest(recipeId uuid.UUID, pictures dto.RecipePicturesUpload, tx *sql.Tx) error {

	updateRequestQuery := fmt.Sprintf(`
		UPDATE %[1]v
		SET pictures=$2
		WHERE recipe_id=$1
	`, recipePicturesUploadsTable)

	if _, err := tx.Exec(updateRequestQuery, recipeId, pictures); err != nil {
		log.Errorf("unable to update recipe %s pictures uploading request: %s", recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}
	return nil
}

func (r *Repository) SetRecipePictures(recipeId uuid.UUID, pictures entity.RecipePictures, version *int32) (int32, error) {
	tx, err := r.startTransaction()
	if err != nil {
		return 0, err
	}

	if err = r.deleteUsedPictureIds(recipeId, pictures, tx); err != nil {
		return 0, err
	}

	var newVersion int32
	dtos := dto.NewRecipePicturesDto(pictures)

	updateRecipeQuery := fmt.Sprintf(`
		UPDATE %[1]v
		SET pictures=$2 version=version+1
		WHERE recipe_id=$1
	`, recipesTable)

	if version != nil {
		updateRecipeQuery += fmt.Sprintf(" AND version=%d", *version)
	}

	updateRecipeQuery += " RETURNING version"

	row := tx.QueryRow(updateRecipeQuery, recipeId, dtos)
	if err = row.Scan(&newVersion); err != nil {
		log.Errorf("unable to set recipe %s pictures: %s", recipeId, err)
		return 0, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return newVersion, commitTransaction(tx)
}

func (r *Repository) deleteUsedPictureIds(recipeId uuid.UUID, pictures entity.RecipePictures, tx *sql.Tx) error {
	picturesUploads, hasRequest, err := r.getExistingRecipePicturesUploadsQuery(recipeId, tx)
	if err != nil {
		return err
	}

	if hasRequest {
		usedPictures := pictures.GetIds()
		var unusedPictures []uuid.UUID
		for _, uploadId := range picturesUploads {
			for _, usedPicture := range usedPictures {
				if usedPicture == uploadId {
					continue
				}
			}
			unusedPictures = append(unusedPictures, uploadId)
		}

		if len(unusedPictures) > 0 {
			if err = r.updateRecipePicturesUploadRequest(recipeId, unusedPictures, tx); err != nil {
				return err
			}
		} else {
			if err = r.deleteRecipePicturesUploadRequest(recipeId, tx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Repository) deleteRecipePicturesUploadRequest(recipeId uuid.UUID, tx *sql.Tx) error {

	deleteRequestQuery := fmt.Sprintf(`
		DELETE FROM %[1]v
		WHERE recipe_id=$1
	`, recipePicturesUploadsTable)

	if _, err := tx.Exec(deleteRequestQuery, recipeId); err != nil {
		log.Errorf("unable to delete recipe %s pictures uploading request: %s", recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}
	return nil
}
