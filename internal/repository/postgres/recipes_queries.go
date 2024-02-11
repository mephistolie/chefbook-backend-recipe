package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres/dto"
	"k8s.io/utils/strings/slices"
	"strconv"
)

var ascendingSortings = []string{entity.SortingTime, entity.SortingCalories}

func (r *Repository) GetRecipes(params entity.RecipesQuery, userId uuid.UUID) []entity.BaseRecipeInfo {
	var recipes []entity.BaseRecipeInfo

	query, args := r.getRecipesByParamsQuery(params, userId)
	log.Debug("Get recipes query generated:\n", query, "\nArgs: ", args)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Errorf("unable to get recipes: %s", err)
		return []entity.BaseRecipeInfo{}
	}

	for rows.Next() {
		var recipe dto.RecipeInfo
		m := pgtype.NewMap()
		if err = rows.Scan(
			&recipe.Id, &recipe.Name,
			&recipe.OwnerId,
			&recipe.Visibility, &recipe.IsEncrypted,
			&recipe.Language, m.SQLScanner(&recipe.Translations),
			&recipe.Rating, &recipe.Votes, &recipe.Score,
			m.SQLScanner(&recipe.Tags), &recipe.Categories, &recipe.IsFavourite, &recipe.IsSaved,
			&recipe.Pictures,
			&recipe.Servings, &recipe.Time,
			&recipe.Calories,
			&recipe.CreationTimestamp, &recipe.UpdateTimestamp, &recipe.Version,
		); err != nil {
			log.Warnf("unable to parse recipe info: ", err)
			continue
		}
		recipes = append(recipes, recipe.Entity(userId))
	}

	return recipes
}

func (r *Repository) getRecipesByParamsQuery(params entity.RecipesQuery, userId uuid.UUID) (string, []interface{}) {
	var args []interface{}

	getRecipesQuery := fmt.Sprintf(`
		SELECT
			%[1]v.recipe_id, %[1]v.name,
			%[1]v.owner_id,
			%[1]v.visibility, %[1]v.encrypted,
			%[1]v.language, %[1]v.translations,
			%[1]v.rating, %[1]v.votes, coalesce(%[3]v.score, 0),
			%[1]v.tags, coalesce(%[2]v.categories, '[]'::jsonb), coalesce(%[2]v.favourite, false),
			(
				SELECT EXISTS
				(
					SELECT 1
					FROM %[2]v
					WHERE %[2]v.recipe_id=%[1]v.recipe_id AND user_id=$1
				)
			) AS saved,
			%[1]v.pictures,
			%[1]v.servings, %[1]v.cooking_time,
			%[1]v.calories,
			%[1]v.creation_timestamp, %[1]v.update_timestamp, %[1]v.version
		FROM
			%[1]v
		LEFT JOIN
			%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id AND %[2]v.user_id=$1
		LEFT JOIN
			%[3]v ON %[3]v.recipe_id=%[1]v.recipe_id AND %[3]v.user_id=$1
	`, recipesTable, usersTable, scoresTable)
	args = append(args, userId)

	whereStatement, whereArgs, argNumber := r.getRecipesWhereStatementByParams(params, 2)
	args = append(args, whereArgs...)

	pagingStatement, pagingArgs := r.getPagingStatement(params, argNumber)
	args = append(args, pagingArgs...)

	return getRecipesQuery + whereStatement + pagingStatement, args
}

func (r *Repository) getRecipesWhereStatementByParams(params entity.RecipesQuery, initArg int) (string, []interface{}, int) {
	var args []interface{}
	argNumber := initArg

	whereStatement := " WHERE"

	if params.Owned && params.Saved {
		whereStatement += fmt.Sprintf(" %s.user_id=$1 AND %s.owner_id=$1", usersTable, recipesTable)
	} else if params.Owned {
		whereStatement += fmt.Sprintf(" %s.owner_id=$1", recipesTable)
	} else if params.Saved {
		if params.AuthorId != nil {
			whereStatement += fmt.Sprintf(" %[1]v.user_id=$1 AND %[2]v.owner_id=%[3]v AND %[2]v.visibility<>'%[4]v'",
				usersTable, recipesTable, *params.AuthorId, model.VisibilityPrivate)
		} else {
			whereStatement += fmt.Sprintf(" %[1]v.user_id=$1 AND (%[2]v.owner_id=$1 OR %[2]v.visibility<>'%[3]v')",
				usersTable, recipesTable, model.VisibilityPrivate)
		}
	} else {
		whereStatement += fmt.Sprintf(" %[1]v.visibility='%[2]v'", recipesTable, model.VisibilityPublic)
	}

	if len(params.RecipeIds) > 0 {
		whereStatement += fmt.Sprintf(" AND %s.recipe_id=ANY($%d)", recipesTable, argNumber)
		args = append(args, params.RecipeIds)
		argNumber += 1
	}

	if !params.Owned && !params.Saved && params.AuthorId != nil {
		whereStatement += fmt.Sprintf(" AND %s.owner_id=$%d", recipesTable, argNumber)
		args = append(args, *params.AuthorId)
		argNumber += 1
	}

	if len(params.Languages) > 0 {
		whereStatement += fmt.Sprintf(" AND (%[1]v.language=ANY($%d) OR %[1]v.translations && $%d)", recipesTable, argNumber)
		args = append(args, params.Languages)
		argNumber += 1
	}

	if params.Search != nil {
		whereStatement += fmt.Sprintf(" AND %s.name LIKE ", recipesTable) + "'%' || $" + strconv.Itoa(argNumber) + " || '%'"
		args = append(args, *params.Search)
		argNumber += 1
	}

	whereStatement += r.getRecipesRangeFilter("rating", params.MinRating, params.MaxRating)
	whereStatement += r.getRecipesRangeFilter("cooking_time", params.MinTime, params.MaxTime)
	whereStatement += r.getRecipesRangeFilter("servings", params.MinServings, params.MaxServings)
	whereStatement += r.getRecipesRangeFilter("calories", params.MinCalories, params.MaxCalories)

	return whereStatement, args, argNumber
}

func (r *Repository) getRecipesRangeFilter(field string, min, max *int32) string {
	filter := ""
	if min != nil {
		filter += fmt.Sprintf(" AND %s.%s>=%d", recipesTable, field, *min)
	}
	if max != nil {
		filter += fmt.Sprintf(" AND %s.%s<=%d", recipesTable, field, *max)
	}
	return filter
}

func (r *Repository) getPagingStatement(params entity.RecipesQuery, initArg int) (string, []interface{}) {
	var args []interface{}
	argNumber := initArg

	pagingStatement := ""

	isAscending := false
	if slices.Contains(ascendingSortings, params.Sorting) {
		isAscending = true
	}

	if params.Sorting == entity.SortingCreationTimestamp && params.LastCreationTimestamp != nil {
		if params.LastRecipeId != nil {
			pagingStatement += fmt.Sprintf(" AND (%[1]v.creation_timestamp, %[1]v.recipe_id) < ($%d, $%d)", recipesTable, argNumber, argNumber+1)
			args = append(args, *params.LastCreationTimestamp, *params.LastRecipeId)
			argNumber += 2
		} else {
			pagingStatement += fmt.Sprintf(" AND %[1]v.creation_timestamp < $%d", recipesTable, argNumber)
			args = append(args, *params.LastCreationTimestamp)
			argNumber += 1
		}
	}
	if params.Sorting == entity.SortingUpdateTimestamp && params.LastUpdateTimestamp != nil {
		if params.LastRecipeId != nil {
			pagingStatement += fmt.Sprintf(" AND (%[1]v.update_timestamp, %[1]v.recipe_id) < ($%d, $%d)", recipesTable, argNumber, argNumber+1)
			args = append(args, *params.LastUpdateTimestamp, *params.LastRecipeId)
			argNumber += 2
		} else {
			pagingStatement += fmt.Sprintf(" AND %[1]v.update_timestamp < $%d", recipesTable, argNumber)
			args = append(args, *params.LastUpdateTimestamp)
			argNumber += 1
		}
	}
	if params.Sorting == entity.SortingRating && params.LastRating != nil {
		if params.LastRecipeId != nil {
			pagingStatement += fmt.Sprintf(" AND (%[1]v.rating, %[1]v.recipe_id) < ($%d, $%d)", recipesTable, argNumber, argNumber+1)
			args = append(args, *params.LastRating, *params.LastRecipeId)
			argNumber += 2
		} else {
			pagingStatement += fmt.Sprintf(" AND %[1]v.rating < $%d", recipesTable, argNumber)
			args = append(args, *params.LastRating)
			argNumber += 1
		}
	}
	if params.Sorting == entity.SortingVotes && params.LastVotes != nil {
		if params.LastRecipeId != nil {
			pagingStatement += fmt.Sprintf(" AND (%[1]v.votes, %[1]v.recipe_id) < ($%d, $%d)", recipesTable, argNumber, argNumber+1)
			args = append(args, *params.LastVotes, *params.LastRecipeId)
			argNumber += 2
		} else {
			pagingStatement += fmt.Sprintf(" AND %[1]v.votes < $%d", recipesTable, argNumber)
			args = append(args, *params.LastVotes)
			argNumber += 1
		}
	}
	if params.Sorting == entity.SortingTime && params.LastTime != nil {
		if params.LastRecipeId != nil {
			pagingStatement += fmt.Sprintf(" AND (%[1]v.cooking_time, %[1]v.recipe_id) > ($%d, $%d)", recipesTable, argNumber, argNumber+1)
			args = append(args, *params.LastTime, *params.LastRecipeId)
			argNumber += 2
		} else {
			pagingStatement += fmt.Sprintf(" AND %[1]v.cooking_time > $%d", recipesTable, argNumber)
			args = append(args, *params.LastTime)
			argNumber += 1
		}
	}
	if params.Sorting == entity.SortingCalories && params.LastCalories != nil {
		if params.LastRecipeId != nil {
			pagingStatement += fmt.Sprintf(" AND (%[1]v.calories, %[1]v.recipe_id) > ($%d, $%d)", recipesTable, argNumber, argNumber+1)
			args = append(args, *params.LastCalories, *params.LastRecipeId)
			argNumber += 2
		} else {
			pagingStatement += fmt.Sprintf(" AND %[1]v.calories > $%d", recipesTable, argNumber)
			args = append(args, *params.LastCalories)
			argNumber += 1
		}
	}

	order := "DESC"
	if isAscending {
		order = "ASC"
	}
	pagingStatement += fmt.Sprintf(" ORDER BY %[1]v %[2]v, recipe_id %[2]v", params.Sorting, order)

	pagingStatement += fmt.Sprintf(" LIMIT %d", params.PageSize)

	return pagingStatement, args
}

func (r *Repository) GetRandomRecipe(userId uuid.UUID, languages *[]string) (entity.BaseRecipe, error) {
	var recipe dto.Recipe

	query := fmt.Sprintf(`
		SELECT
			%[1]v.recipe_id, %[1]v.name,
			%[1]v.owner_id,
			%[1]v.visibility, %[1]v.encrypted,
			%[1]v.language, %[1]v.description,
			%[1]v.rating, %[1]v.votes, coalesce(%[3]v.score, 0),
			%[1]v.tags,
			%[1]v.ingredients, %[1]v.cooking, %[1]v.pictures,
			%[1]v.servings, %[1]v.cooking_time,
			%[1]v.calories, %[1]v.protein, %[1]v.fats, %[1]v.carbohydrates,
			%[1]v.creation_timestamp, %[1]v.update_timestamp, %[1]v.version
		FROM
			%[1]v
		LEFT JOIN
			%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id AND %[2]v.user_id=$1
		LEFT JOIN
			%[3]v ON %[3]v.recipe_id=%[1]v.recipe_id AND %[3]v.user_id=$1
		WHERE
			%[1]v.visibility='%[4]v' AND %[1]v.owner_id<>$1 AND
			NOT EXISTS
			(
				SELECT 1
				FROM %[2]v
				WHERE %[2]v.recipe_id=%[1]v.recipe_id AND user_id=$1
			)
	`, recipesTable, usersTable, scoresTable, model.VisibilityPublic)

	var args []interface{}
	args = append(args, userId)

	if languages != nil {
		query += fmt.Sprintf(" AND (%[1]v.language=ANY($2) OR %[1]v.translations && $2)", recipesTable)
		args = append(args, *languages)
	}

	query += " ORDER BY RANDOM() LIMIT 1"

	row := r.db.QueryRow(query, args...)
	m := pgtype.NewMap()
	if err := row.Scan(
		&recipe.Id, &recipe.Name,
		&recipe.OwnerId,
		&recipe.Visibility, &recipe.IsEncrypted,
		&recipe.Language, &recipe.Description,
		&recipe.Rating, &recipe.Votes, &recipe.Score,
		m.SQLScanner(&recipe.Tags),
		&recipe.Ingredients, &recipe.Cooking, &recipe.Pictures,
		&recipe.Servings, &recipe.Time,
		&recipe.Calories, &recipe.Protein, &recipe.Fats, &recipe.Carbohydrates,
		&recipe.CreationTimestamp, &recipe.UpdateTimestamp, &recipe.Version,
	); err != nil {
		log.Debug("unable to get random recipe for user %s: %s", userId, err)
		return entity.BaseRecipe{}, fail.GrpcNotFound
	}

	recipe.Translations, _ = r.GetRecipeTranslations(recipe.Id)
	delete(recipe.Translations, recipe.Language)

	return recipe.Entity(userId), nil
}

func (r *Repository) GetRecipeNames(recipeIds []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]string, error) {
	recipeNames := make(map[uuid.UUID]string)

	query := fmt.Sprintf(`
		SELECT recipe_id, name
		FROM %[1]v
		WHERE recipe_id=ANY($1) AND (owner_id=$2 OR visibility<>'%[2]v')
	`, recipesTable, model.VisibilityPrivate)

	rows, err := r.db.Query(query, recipeIds, userId)
	if err != nil {
		log.Errorf("unable to get recipe names: %s", err)
		return map[uuid.UUID]string{}, fail.GrpcUnknown
	}

	for rows.Next() {
		var id uuid.UUID
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			log.Warnf("unable to parse recipe name: ", err)
			continue
		}
		recipeNames[id] = name
	}

	return recipeNames, nil
}
