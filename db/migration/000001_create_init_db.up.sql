CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "hashed_password" bytea NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "preparation_time_in_mins" int NOT NULL,
  "difficulty_level" int NOT NULL,
  "cuisine_type" varchar(255) NOT NULL,
  "calorie_count_per_serving" int NOT NULL,
  "servings_count" int NOT NULL,
  "preparation_steps" text NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ingredients" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "calorie_count_per_unit" int NOT NULL
);

CREATE TABLE "recipe_ingredients" (
  "id" bigserial PRIMARY KEY,
  "recipe_id" int NOT NULL,
  "ingredient_id" int NOT NULL,
  "quantity" float NOT NULL,
  "unit" varchar(255) NOT NULL
);

ALTER TABLE "recipes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "recipe_ingredients" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");

ALTER TABLE "recipe_ingredients" ADD FOREIGN KEY ("ingredient_id") REFERENCES "ingredients" ("id");