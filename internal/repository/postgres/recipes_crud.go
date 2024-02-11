package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/mq/model"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/mq"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres/dto"
)

func (r *Repository) CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error) {
	var id uuid.UUID
	if input.RecipeId != nil {
		id = *input.RecipeId
	} else {
		id = uuid.New()
	}

	tx, err := r.startTransaction()
	if err != nil {
		return uuid.UUID{}, 0, err
	}

	createRecipeQuery := fmt.Sprintf(`
		INSERT INTO %s
			(
				recipe_id, name,
				owner_id,
				visibility, encrypted,
				language, description,
				tags,
				ingredients, cooking,
				servings, cooking_time,
				calories, protein, fats, carbohydrates
			)
		VALUES
			(
				$1, $2,
				$3,
				$4, $5,
				$6, $7,
				$8,
				$9, $10,
				$11, $12,
				$13, $14, $15, $16
			)
	`, recipesTable)

	macronutrients := entity.Macronutrients{}
	if input.Macronutrients != nil {
		macronutrients = *input.Macronutrients
	}

	if _, err = tx.Exec(createRecipeQuery,
		id, input.Name,
		input.UserId,
		input.Visibility, input.IsEncrypted,
		input.Language, input.Description,
		append([]string{}, input.Tags...),
		dto.NewIngredients(input.Ingredients), dto.NewCooking(input.Cooking),
		input.Servings, input.Time,
		input.Calories, macronutrients.Protein, macronutrients.Fats, macronutrients.Carbohydrates,
	); err != nil {
		if isUniqueViolationError(err) {
			return uuid.UUID{}, 0, recipeFail.GrpcRecipeExists
		}
		log.Errorf("unable to create recipe: %s", err)
		return uuid.UUID{}, 0, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	if input.CreationTimestamp != nil {
		setCreationTimestampQuery := fmt.Sprintf(`
			UPDATE %s
			SET creation_timestamp=$2
			WHERE recipe_id=$1
		`, recipesTable)

		if _, err = tx.Exec(setCreationTimestampQuery, id, *input.CreationTimestamp); err != nil {
			log.Error("unable to set recipe creation timestamp: ", err)
			return uuid.UUID{}, 0, errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}
	}

	addToRecipeBookQuery := fmt.Sprintf(`
			INSERT INTO %s (recipe_id, user_id)
			VALUES ($1, $2)
		`, usersTable)

	if _, err = tx.Exec(addToRecipeBookQuery, id, input.UserId); err != nil {
		log.Errorf("unable to add recipe to owner %s recipe book: %s", input.UserId, err)
		return uuid.UUID{}, 0, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return id, 1, commitTransaction(tx)
}

func (r *Repository) GetRecipe(recipeId, userId uuid.UUID) (entity.BaseRecipe, error) {
	var recipe dto.Recipe

	query := fmt.Sprintf(`
		SELECT
			%[1]v.recipe_id, %[1]v.name,
			%[1]v.owner_id,
			%[1]v.visibility, %[1]v.encrypted,
			%[1]v.language, %[1]v.description,
			%[1]v.rating, %[1]v.votes, coalesce(%[3]v.score, 0),
			%[1]v.tags, coalesce(%[2]v.categories, '[]'::jsonb), coalesce(%[2]v.favourite, false),
			(
				SELECT EXISTS
				(
					SELECT 1
					FROM %[2]v
					WHERE %[2]v.recipe_id=%[1]v.recipe_id AND user_id=$2
				)
			) AS saved,
			%[1]v.ingredients, %[1]v.cooking, %[1]v.pictures,
			%[1]v.servings, %[1]v.cooking_time,
			%[1]v.calories, %[1]v.protein, %[1]v.fats, %[1]v.carbohydrates,
			%[1]v.creation_timestamp, %[1]v.update_timestamp, %[1]v.version
		FROM
			%[1]v
		LEFT JOIN
			%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id AND %[2]v.user_id=$2
		LEFT JOIN
			%[3]v ON %[3]v.recipe_id=%[1]v.recipe_id AND %[3]v.user_id=$2
		WHERE %[1]v.recipe_id=$1
	`, recipesTable, usersTable, scoresTable)

	row := r.db.QueryRow(query, recipeId, userId)
	m := pgtype.NewMap()
	if err := row.Scan(
		&recipe.Id, &recipe.Name,
		&recipe.OwnerId,
		&recipe.Visibility, &recipe.IsEncrypted,
		&recipe.Language, &recipe.Description,
		&recipe.Rating, &recipe.Votes, &recipe.Score,
		m.SQLScanner(&recipe.Tags), &recipe.Categories, &recipe.IsFavourite, &recipe.IsSaved,
		&recipe.Ingredients, &recipe.Cooking, &recipe.Pictures,
		&recipe.Servings, &recipe.Time,
		&recipe.Calories, &recipe.Protein, &recipe.Fats, &recipe.Carbohydrates,
		&recipe.CreationTimestamp, &recipe.UpdateTimestamp, &recipe.Version,
	); err != nil {
		log.Warnf("unable to get recipe %s for user %s: %s", recipeId, userId, err)
		return entity.BaseRecipe{}, fail.GrpcNotFound
	}

	recipe.Translations, _ = r.GetRecipeTranslations(recipeId)
	delete(recipe.Translations, recipe.Language)

	return recipe.Entity(userId), nil
}

func (r *Repository) UpdateRecipe(input entity.RecipeInput) (int32, error) {
	var version int32

	query := fmt.Sprintf(`
		UPDATE %[1]v
		SET 
			name=$1,
			visibility=$2, encrypted=$3,
			language=$4, description=$5,
			tags=$6,
			ingredients=$7, cooking=$8,
			servings=$9, cooking_time=$10,
			calories=$11, protein=$12, fats=$13, carbohydrates=$14,
			version=version+1, update_timestamp=now()::timestamp
		WHERE recipe_id=$15
	`, recipesTable)

	macronutrients := entity.Macronutrients{}
	if input.Macronutrients != nil {
		macronutrients = *input.Macronutrients
	}

	args := []interface{}{
		input.Name,
		input.Visibility, input.IsEncrypted,
		input.Language, input.Description,
		append([]string{}, input.Tags...),
		dto.NewIngredients(input.Ingredients), dto.NewCooking(input.Cooking),
		input.Servings, input.Time,
		input.Calories, macronutrients.Protein, macronutrients.Fats, macronutrients.Carbohydrates,
		*input.RecipeId,
	}

	if input.Version != nil {
		query += " AND version=$16"
		args = append(args, *input.Version)
	}

	query += " RETURNING version"

	if err := r.db.Get(&version, query, args...); err != nil {
		if input.Version != nil {
			log.Warnf("try to update recipe %s with outdated version %d: %s", *input.RecipeId, *input.Version, err)
			return 0, recipeFail.GrpcOutdatedVersion
		} else {
			log.Errorf("unable to update recipe %s: %s", *input.RecipeId, err)
			return 0, fail.GrpcUnknown
		}
	}

	return version, nil
}

func (r *Repository) SetRecipeTags(recipeId uuid.UUID, tags []string) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET tags=$2
		WHERE recipe_id=$1
	`, recipesTable)

	if _, err := r.db.Exec(query, recipeId, tags); err != nil {
		log.Warnf("unable to set recipe %s tags: %s", recipeId, err)
		return fail.GrpcNotFound
	}
	return nil
}

func (r *Repository) DeleteRecipe(recipeId uuid.UUID) (*model.MessageData, error) {
	tx, err := r.startTransaction()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE recipe_id=$1
	`, recipesTable)

	if _, err = tx.Exec(query, recipeId); err != nil {
		log.Errorf("unable to delete recipe %s: %s", recipeId, err)
		return nil, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	msg, err := r.addRecipeDeletedMsg(recipeId, tx)
	if err != nil {
		return nil, err
	}

	return msg, commitTransaction(tx)
}

func (r *Repository) addRecipeDeletedMsg(recipeId uuid.UUID, tx *sql.Tx) (*model.MessageData, error) {
	msgBody := api.MsgBodyRecipeDeleted{RecipeId: recipeId}
	msgBodyBson, err := json.Marshal(msgBody)
	if err != nil {
		log.Error("unable to marshal recipe deleted message body: ", err)
		return nil, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}
	msgInfo := model.MessageData{
		Id:       uuid.New(),
		Exchange: api.ExchangeRecipes,
		Type:     api.MsgTypeRecipeDeleted,
		Body:     msgBodyBson,
	}

	return &msgInfo, r.createOutboxMsg(&msgInfo, tx)
}
