-- name: CreateRecipe :one
INSERT INTO recipes (
    name, preparation_time_in_mins, difficulty_level, 
    cuisine_type_id, calorie_count_per_serving,
    servings_count, preparation_steps, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING *;

-- name: GetRecipeByID :one
SELECT * FROM recipes 
WHERE id = $1;
