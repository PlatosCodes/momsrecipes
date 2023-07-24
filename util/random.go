package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomRecipeRequest() RandomRecipeParams {
	return RandomRecipeParams{
		Name:                   RandomString(6),
		PreparationTimeInMins:  RandomInt(1, 100),
		DifficultyLevel:        RandomInt(1, 10),
		CuisineType:            RandomString(5),
		CalorieCountPerServing: RandomInt(200, 1000),
		ServingsCount:          RandomInt(1, 5),
		PreparationSteps:       RandomString(100),
		UserID:                 1,
	}
}

type RandomRecipeParams struct {
	Name                   string `json:"name"`
	PreparationTimeInMins  int32  `json:"preparation_time_in_mins"`
	DifficultyLevel        int32  `json:"difficulty_level"`
	CuisineType            string `json:"cuisine_type"`
	CalorieCountPerServing int32  `json:"calorie_count_per_serving"`
	ServingsCount          int32  `json:"servings_count"`
	PreparationSteps       string `json:"preparation_steps"`
	UserID                 int32  `json:"user_id"`
}
