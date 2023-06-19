package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres/dto"
)

func (r *Repository) CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error) {
	var id uuid.UUID
	if input.Id != nil {
		id = *input.Id
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
		dto.NewTags(input.Tags),
		dto.NewIngredients(input.Ingredients), dto.NewCooking(input.Cooking),
		input.Servings, input.Time,
		input.Calories, macronutrients.Protein, macronutrients.Fats, macronutrients.Carbohydrates,
	); err != nil {
		log.Errorf("unable to create recipe: %s", err)
		return uuid.UUID{}, 0, errorWithTransactionRollback(tx, err)
	}

	if input.CreationTimestamp != nil {
		setCreationTimestampQuery := fmt.Sprintf(`
			UPDATE %s
			SET creation_timestamp=$2
			WHERE recipe_id=$1
		`, recipesTable)

		if _, err = tx.Exec(setCreationTimestampQuery, id, *input.CreationTimestamp); err != nil {
			log.Error("unable to set recipe creation timestamp: ", err)
			return uuid.UUID{}, 0, errorWithTransactionRollback(tx, err)
		}
	}

	addToRecipeBookQuery := fmt.Sprintf(`
			INSERT INTO %s (recipe_id, user_id)
			VALUES ($1, $2)
		`, usersTable)

	if _, err = tx.Exec(addToRecipeBookQuery, id, input.UserId); err != nil {
		log.Errorf("unable to add recipe to owner %s recipe book: %s", input.UserId, err)
		return uuid.UUID{}, 0, errorWithTransactionRollback(tx, err)
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
			%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id
		LEFT JOIN
			%[3]v ON %[3]v.recipe_id=%[1]v.recipe_id
		WHERE %[1]v.recipe_id=$1
	`, recipesTable, usersTable, scoresTable)

	row := r.db.QueryRow(query, recipeId)
	if err := row.Scan(&recipe); err != nil {
		log.Warnf("unable to get recipe %s for user %s; %s", recipeId, userId, err)
		return entity.BaseRecipe{}, fail.GrpcNotFound
	}

	return recipe.Entity(userId), nil
}

func (r *Repository) UpdateRecipe(input entity.RecipeInput) (int32, error) {
	var version int32

	query := fmt.Sprintf(`
		UPDATE %s
		SET 
			name=$1
			visibility=$2, encrypted=$3,
			language=$4, description=$5,
			tags=$6,
			ingredients=$7, cooking=$8,
			servings=$9, cooking_time=$10,
			calories=$11, protein=$12, fats=$13, carbohydrates=$14,
			version=version+1
		WHERE recipe_id=$15 AND owner_id=$16
	`, recipesTable)

	macronutrients := entity.Macronutrients{}
	if input.Macronutrients != nil {
		macronutrients = *input.Macronutrients
	}

	args := []interface{}{
		input.Name,
		input.UserId,
		input.Visibility, input.IsEncrypted,
		input.Language, input.Description,
		input.Tags,
		dto.NewIngredients(input.Ingredients), dto.NewCooking(input.Cooking),
		input.Servings, input.Time,
		input.Calories, macronutrients.Protein, macronutrients.Fats, macronutrients.Carbohydrates,
		*input.Id, input.UserId,
	}

	if input.Version != nil {
		query += " AND version=$17"
		args = append(args, *input.Version)
	}

	query += " RETURNING version"

	if err := r.db.Get(&version, query, args...); err != nil {
		log.Errorf("unable to update recipe %s: %s", *input.Id, err)
		return 0, fail.GrpcUnknown
	}

	return version, nil
}

func (r *Repository) DeleteRecipe(recipeId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE recipe_id=$1 AND owner_id=$2
	`, recipesTable)

	if _, err := r.db.Exec(query, recipeId, userId); err != nil {
		log.Errorf("unable to delete recipe %s: %s", recipeId, err)
		return fail.GrpcUnknown
	}

	return nil
}
