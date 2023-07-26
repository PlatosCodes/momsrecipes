package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var randSeed struct {
	value int64
	once  sync.Once
}

var randGenerator *rand.Rand

func Rand() *rand.Rand {
	randSeed.once.Do(func() {
		randSeed.value = time.Now().UnixMicro()
		randGenerator = rand.New(rand.NewSource(randSeed.value))
	})
	return randGenerator
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[Rand().Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(10) + strconv.Itoa(time.Now().Nanosecond())

}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomRecipeRequest() RandomRecipeParams {
	return RandomRecipeParams{
		Name:                   RandomString(6),
		PreparationTimeInMins:  Rand().Int31(),
		DifficultyLevel:        Rand().Int31(),
		CuisineType:            RandomString(5),
		CalorieCountPerServing: Rand().Int31(),
		ServingsCount:          Rand().Int31(),
		PreparationSteps:       RandomString(100),
		UserID:                 Rand().Int63(),
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
	UserID                 int64  `json:"user_id"`
}
