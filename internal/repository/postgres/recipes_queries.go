package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
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
		recipe := dto.RecipeInfo{}
		if err = rows.Scan(
			&recipe.Id, &recipe.Name,
			&recipe.OwnerId,
			&recipe.Visibility, &recipe.IsEncrypted,
			&recipe.Language,
			&recipe.Rating, &recipe.Votes, &recipe.Score,
			&recipe.Tags, &recipe.Categories, &recipe.IsFavourite, &recipe.IsSaved,
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
			%[1]v.language, 
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
			%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id
		LEFT JOIN
			%[3]v ON %[3]v.recipe_id=%[1]v.recipe_id
	`, recipesTable, usersTable, scoresTable)
	args = append(args, userId)

	whereStatement, whereArgs, argNumber := r.getRecipesWhereStatementByParams(params, 1)
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
				usersTable, recipesTable, *params.AuthorId, entity.VisibilityPrivate)
		} else {
			whereStatement += fmt.Sprintf(" %[1]v.user_id=$1 AND (%[2]v.owner_id=$1 OR %[2]v.visibility<>'%[3]v')",
				usersTable, recipesTable, entity.VisibilityPrivate)
		}
	} else {
		whereStatement += fmt.Sprintf(" %[1]v.visibility='%[2]v'", recipesTable, entity.VisibilityPublic)
	}

	if !params.Owned && !params.Saved && params.AuthorId != nil {
		whereStatement += fmt.Sprintf(" AND %s.owner_id=%s", recipesTable, *params.AuthorId)
	}

	if params.Languages != nil && len(*params.Languages) > 0 {
		whereStatement += fmt.Sprintf(" AND %s.language=ANY($%d)", recipesTable, argNumber)
		args = append(args, *params.Languages)
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
		filter += fmt.Sprintf(" AND %s.%s>=%d", recipesTable, field, min)
	}
	if max != nil {
		filter += fmt.Sprintf(" AND %s.%s<=%d", recipesTable, field, max)
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
		pagingStatement += fmt.Sprintf(" AND %s.creation_timestamp<%s", recipesTable, *params.LastCreationTimestamp)
	}
	if params.Sorting == entity.SortingUpdateTimestamp && params.LastUpdateTimestamp != nil {
		pagingStatement += fmt.Sprintf(" AND %s.update_timestamp<%s", recipesTable, *params.LastUpdateTimestamp)
	}
	if params.Sorting == entity.SortingRating && params.LastRating != nil {
		pagingStatement += fmt.Sprintf(" AND %s.rating<=$%d", recipesTable, argNumber)
		args = append(args, *params.LastRating)
		argNumber += 1
	}
	if params.Sorting == entity.SortingVotes && params.LastVotes != nil {
		pagingStatement += fmt.Sprintf(" AND %s.rating<=$%d", recipesTable, *params.LastVotes)
	}
	if params.Sorting == entity.SortingTime && params.LastTime != nil {
		pagingStatement += fmt.Sprintf(" AND %s.cooking_time>=$%d", recipesTable, *params.LastTime)
	}
	if params.Sorting == entity.SortingCalories && params.LastCalories != nil {
		pagingStatement += fmt.Sprintf(" AND %s.calories>=$%d", recipesTable, *params.LastCalories)
	}

	if params.LastRecipeId != nil {
		if isAscending {
			pagingStatement += fmt.Sprintf(" AND %s.recipe_id>$%d", recipesTable, *params.LastRecipeId)
		} else {
			pagingStatement += fmt.Sprintf(" AND %s.recipe_id<$%d", recipesTable, *params.LastRecipeId)
		}
	}

	order := "DESC"
	if isAscending {
		order += "ASC"
	}
	pagingStatement += fmt.Sprintf(" ORDER BY %[1]v %[2]v, recipe_id %[2]v", params.Sorting, order)

	if params.PageSize != nil {
		pagingStatement += fmt.Sprintf(" LIMIT %d", *params.PageSize)
	}

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
			%[1]v.tags, coalesce(%[2]v.categories, '[]'::jsonb), coalesce(%[2]v.favourite, false),
			(
				SELECT EXISTS
				(
					SELECT 1
					FROM %[2]v
					WHERE %[2]v.recipe_id=%[1]v.recipe_id AND user_id=$1
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
		WHERE
			%[1]v.visibility='%[4]v' AND saved=false AND owner_id<>$1
	`, recipesTable, usersTable, scoresTable, entity.VisibilityPublic)

	var args []interface{}
	args = append(args, userId)

	if languages != nil {
		query += fmt.Sprintf(" AND %s.language=ANY($2)", recipesTable)
		args = append(args, *languages)
	}

	query += " ORDER BY RANDOM() LIMIT 1"

	row := r.db.QueryRow(query, args...)
	if err := row.Scan(
		&recipe.Id, &recipe.Name,
		&recipe.OwnerId,
		&recipe.Visibility, &recipe.IsEncrypted,
		&recipe.Language, &recipe.Description,
		&recipe.Rating, &recipe.Votes, &recipe.Score,
		&recipe.Tags, &recipe.Categories, &recipe.IsFavourite, &recipe.IsSaved,
		&recipe.Ingredients, &recipe.Cooking, &recipe.Pictures,
		&recipe.Servings, &recipe.Time,
		&recipe.Calories, &recipe.Protein, &recipe.Fats, &recipe.Carbohydrates,
		&recipe.CreationTimestamp, &recipe.UpdateTimestamp, &recipe.Version,
	); err != nil {
		log.Warnf("unable to get random recipe for user %s: %s", userId, err)
		return entity.BaseRecipe{}, fail.GrpcNotFound
	}

	return recipe.Entity(userId), nil
}

func (r *Repository) GetRecipeBook(userId uuid.UUID) ([]entity.BaseRecipeState, error) {
	var recipes []entity.BaseRecipeState

	query := fmt.Sprintf(`
		SELECT
			%[1]v.recipe_id,
			%[1]v.owner_id,
			%[1]v.rating, %[1]v.votes, coalesce(%[3]v.score, 0),
			%[1]v.tags, coalesce(%[2]v.categories, '[]'::jsonb), coalesce(%[2]v.favourite, false),
			%[1]v.version
		FROM
			%[1]v
		LEFT JOIN
			%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id
		LEFT JOIN
			%[3]v ON %[3]v.recipe_id=%[1]v.recipe_id
		WHERE
			%[2]v.user_id=$1 AND (%[1]v.owner_id=$1 OR %[1]v.visibility<>'%[4]v')
	`, recipesTable, usersTable, scoresTable, entity.VisibilityPrivate)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		log.Errorf("unable to get recipes: %s", err)
		return []entity.BaseRecipeState{}, fail.GrpcUnknown
	}

	for rows.Next() {
		recipe := dto.RecipeState{}
		if err = rows.Scan(
			&recipe.Id,
			&recipe.OwnerId,
			&recipe.Rating, &recipe.Votes, &recipe.Score,
			&recipe.Tags, &recipe.Categories, &recipe.IsFavourite,
			&recipe.Version,
		); err != nil {
			log.Warnf("unable to parse recipe info: ", err)
			continue
		}
		recipes = append(recipes, recipe.Entity())
	}

	return recipes, nil
}

func (r *Repository) GetRecipeNames(recipeIds []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]string, error) {
	recipeNames := make(map[uuid.UUID]string)

	query := fmt.Sprintf(`
		SELECT recipe_id, name
		FROM %[1]v
		WHERE recipe_id=ANY($1) AND (owner_id=$2 OR visibility<>'%[2]v')
	`, recipesTable, entity.VisibilityPrivate)

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
