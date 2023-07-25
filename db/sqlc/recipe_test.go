package db

import (
	"context"
	"testing"
	"time"

	"github.com/PlatosCodes/momsrecipes/util"
	"github.com/stretchr/testify/require"
)

func createRandomRecipe(t *testing.T) Recipe {
	user := createRandomUser(t)
	arg := randomRecipeParams(t, user.ID)

	recipe, err := testQueries.CreateRecipe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipe)

	require.Equal(t, arg.Name, recipe.Name)
	require.Equal(t, arg.PreparationTimeInMins, recipe.PreparationTimeInMins)
	require.Equal(t, arg.CalorieCountPerServing, recipe.CalorieCountPerServing)
	require.Equal(t, arg.CuisineType, recipe.CuisineType)
	require.Equal(t, arg.DifficultyLevel, recipe.DifficultyLevel)
	require.Equal(t, arg.PreparationSteps, recipe.PreparationSteps)
	require.Equal(t, arg.ServingsCount, recipe.ServingsCount)

	require.NotZero(t, recipe.ID)
	require.NotZero(t, recipe.CreatedAt)

	return recipe
}

func TestCreateRecipe(t *testing.T) {
	createRandomRecipe(t)
}

func TestGetRecipe(t *testing.T) {
	recipe1 := createRandomRecipe(t)
	recipe2, err := testQueries.GetRecipeByID(context.Background(), recipe1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, recipe2)

	require.Equal(t, recipe1.ID, recipe2.ID)
	require.Equal(t, recipe1.Name, recipe2.Name)
	require.Equal(t, recipe1.CalorieCountPerServing, recipe2.CalorieCountPerServing)
	require.Equal(t, recipe1.CuisineType, recipe2.CuisineType)
	require.Equal(t, recipe1.DifficultyLevel, recipe2.DifficultyLevel)
	require.Equal(t, recipe1.PreparationSteps, recipe2.PreparationSteps)
	require.Equal(t, recipe1.ServingsCount, recipe2.ServingsCount)
	require.Equal(t, recipe1.PreparationTimeInMins, recipe2.PreparationTimeInMins)

	require.WithinDuration(t, recipe1.CreatedAt, recipe2.CreatedAt, time.Second)
}

func randomRecipeParams(t *testing.T, userID int64) CreateRecipeParams {
	return CreateRecipeParams{
		Name:                   util.RandomString(6),
		PreparationTimeInMins:  util.RandomInt(1, 100),
		DifficultyLevel:        util.RandomInt(1, 10),
		CuisineType:            util.RandomString(5),
		CalorieCountPerServing: util.RandomInt(200, 1000),
		ServingsCount:          util.RandomInt(1, 5),
		PreparationSteps:       util.RandomString(100),
		UserID:                 userID,
	}
}
