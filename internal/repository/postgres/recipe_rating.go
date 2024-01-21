package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/mq/model"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/mq"
)

func (r *Repository) GetRecipeRatingAndVotes(recipeId uuid.UUID) (float32, int, error) {
	var rating float32
	var votes int

	query := fmt.Sprintf(`
		SELECT rating, votes
		FROM %s
		WHERE recipe_id=$1
	`, recipesTable)

	row := r.db.QueryRow(query, recipeId)
	if err := row.Scan(&rating, &votes); err != nil {
		log.Warnf("unable to parse rating and votes for recipe %s: %s", recipeId, err)
		return 0, 0, fail.GrpcNotFound
	}

	return rating, votes, nil
}

func (r *Repository) GetUserRecipeScore(recipeId, userId uuid.UUID) int {
	var score int

	query := fmt.Sprintf(`
		SELECT score
		FROM %s
		WHERE recipe_id=$1 AND user_id=$2
	`, scoresTable)

	row := r.db.QueryRow(query, recipeId, userId)
	if err := row.Scan(&score); err != nil {
		return 0
	}

	return score
}

func (r *Repository) RateRecipe(recipeId, userId uuid.UUID, score int) (*model.MessageData, error) {
	previousScore := r.GetUserRecipeScore(recipeId, userId)

	scoreDiff := score - previousScore
	if scoreDiff == 0 {
		return nil, nil
	}

	tx, err := r.startTransaction()
	if err != nil {
		return nil, err
	}

	if previousScore == 0 {
		err = r.addUserScore(tx, recipeId, userId, scoreDiff)
	} else if score == 0 {
		err = r.deleteUserScore(tx, recipeId, userId, scoreDiff)
	} else {
		err = r.changeUserScore(tx, recipeId, userId, scoreDiff)
	}
	if err != nil {
		return nil, err
	}

	msg, err := r.addRecipeRatingChangedMsg(recipeId, userId, scoreDiff, tx)
	if err != nil {
		return nil, err
	}

	return msg, commitTransaction(tx)
}

func (r *Repository) addUserScore(tx *sql.Tx, recipeId, userId uuid.UUID, score int) error {
	addScoreQuery := fmt.Sprintf(`
		INSERT INTO %s (recipe_id, user_id, score)
		VALUES ($1, $2, $3)
	`, scoresTable)

	if _, err := tx.Query(addScoreQuery, recipeId, userId, score); err != nil {
		log.Errorf("unable to add user %s score for recipe %s: %s", userId, recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	updateRatingQuery := fmt.Sprintf(`
		UPDATE %s
		SET
			rating=(ceil(rating*votes)+$1)/(votes+1),
			votes=votes+1
		WHERE recipe_id=$2
	`, recipesTable)

	if _, err := tx.Query(updateRatingQuery, score, recipeId); err != nil {
		log.Errorf("unable to update rating for recipe %s: %s", recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return nil
}

func (r *Repository) changeUserScore(tx *sql.Tx, recipeId, userId uuid.UUID, scoreDiff int) error {
	changeScoreQuery := fmt.Sprintf(`
		UPDATE %s
		SET score=score+$1
		WHERE recipe_id=$2 AND user_id=$3
	`, scoresTable)

	if _, err := tx.Query(changeScoreQuery, scoreDiff, recipeId, userId); err != nil {
		log.Errorf("unable to change user %s score for recipe %s: %s", userId, recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	updateRatingQuery := fmt.Sprintf(`
		UPDATE %s
		SET rating=(ceil(rating*votes)+$1)/votes
		WHERE recipe_id=$2
	`, recipesTable)

	if _, err := tx.Query(updateRatingQuery, scoreDiff, recipeId); err != nil {
		log.Errorf("unable to update rating for recipe %s: %s", recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return nil
}

func (r *Repository) deleteUserScore(tx *sql.Tx, recipeId, userId uuid.UUID, scoreDiff int) error {
	changeScoreQuery := fmt.Sprintf(`
		DELETE FROM %s
		WHERE recipe_id=$1 AND user_id=$2
	`, scoresTable)

	if _, err := tx.Query(changeScoreQuery, recipeId, userId); err != nil {
		log.Errorf("unable to delete user %s score for recipe %s: %s", userId, recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	updateRatingQuery := fmt.Sprintf(`
		UPDATE %s
		SET
			rating=(ceil(rating*votes)+$1)/greatest(votes-1, 1),
			votes=votes-1
		WHERE recipe_id=$2
	`, recipesTable)

	if _, err := tx.Query(updateRatingQuery, scoreDiff, recipeId); err != nil {
		log.Errorf("unable to update rating for recipe %s: %s", recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return nil
}

func (r *Repository) addRecipeRatingChangedMsg(recipeId, userId uuid.UUID, scoreDiff int, tx *sql.Tx) (*model.MessageData, error) {
	var ownerId uuid.UUID

	getOwnerIdQuery := fmt.Sprintf(`
		SELECT owner_id
		FROM %s
		WHERE recipe_id=$1
	`, recipesTable)

	row := tx.QueryRow(getOwnerIdQuery, recipeId)
	if err := row.Scan(&ownerId); err != nil {
		log.Warnf("unable to get recipe %s owner: %s", recipeId, err)
		return nil, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	msgBody := api.MsgBodyRecipeRatingChanged{
		RecipeId:  recipeId,
		OwnerId:   ownerId,
		ScoreDiff: scoreDiff,
		UserId:    userId,
	}
	msgBodyBson, err := json.Marshal(msgBody)
	if err != nil {
		log.Error("unable to marshal recipe rating changed message body: ", err)
		return nil, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}
	msgInfo := model.MessageData{
		Id:       uuid.New(),
		Exchange: api.ExchangeRecipes,
		Type:     api.MsgTypeRecipeRatingChanged,
		Body:     msgBodyBson,
	}

	return &msgInfo, r.createOutboxMsg(&msgInfo, tx)
}
