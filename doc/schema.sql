-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-07-24T13:22:00.368Z

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "hashed_password" varchar(255) NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT (0001-01-01 00:00:00),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "preparation_time_in_mins" int NOT NULL,
  "difficulty_level" int NOT NULL,
  "cuisine_type_id" int NOT NULL,
  "calorie_count_per_serving" int NOT NULL,
  "servings_count" int NOT NULL,
  "preparation_steps" text NOT NULL,
  "user_id" int NOT NULL
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

CREATE TABLE "cuisine_types" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL
);

ALTER TABLE "recipes" ADD FOREIGN KEY ("cuisine_type_id") REFERENCES "cuisine_types" ("id");

ALTER TABLE "recipes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "recipe_ingredients" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");

ALTER TABLE "recipe_ingredients" ADD FOREIGN KEY ("ingredient_id") REFERENCES "ingredients" ("id");