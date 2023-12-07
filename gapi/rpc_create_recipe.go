package gapi

import (
	"context"
	"log"

	db "github.com/PlatosCodes/momsrecipes/db/sqlc"
	"github.com/PlatosCodes/momsrecipes/pb"
	"github.com/PlatosCodes/momsrecipes/val"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateRecipe(ctx context.Context, req *pb.CreateRecipeRequest) (*pb.CreateRecipeResponse, error) {
	// Verify the access token
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateCreateRecipeRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// Check if the user ID in the request matches the user ID in the token
	if authPayload.UserID != req.GetUserId() {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized user")
	}

	arg := db.CreateRecipeParams{
		Name:                   req.GetName(),
		PreparationTimeInMins:  req.GetPreparationTimeInMins(),
		DifficultyLevel:        req.GetDifficultyLevel(),
		CuisineType:            req.GetCuisineType(),
		CalorieCountPerServing: req.GetCalorieCountPerServing(),
		ServingsCount:          req.GetServingsCount(),
		PreparationSteps:       req.GetPreparationSteps(),
		UserID:                 req.GetUserId(),
	}

	recipe, err := server.Store.CreateRecipe(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println("pqErr", err)
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	return &pb.CreateRecipeResponse{
		Recipe: &pb.Recipe{
			Id:                     recipe.ID,
			Name:                   recipe.Name,
			PreparationTimeInMins:  recipe.PreparationTimeInMins,
			DifficultyLevel:        recipe.DifficultyLevel,
			CuisineType:            recipe.CuisineType,
			CalorieCountPerServing: recipe.CalorieCountPerServing,
			ServingsCount:          recipe.ServingsCount,
			PreparationSteps:       recipe.PreparationSteps,
			UserId:                 recipe.UserID,
		},
	}, nil
}

func validateCreateRecipeRequest(req *pb.CreateRecipeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateRecipeName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := val.ValidatePreparationTime(req.GetPreparationTimeInMins()); err != nil {
		violations = append(violations, fieldViolation("preparation_time_in_mins", err))
	}
	if err := val.ValidateDifficultyLevel(req.GetDifficultyLevel()); err != nil {
		violations = append(violations, fieldViolation("difficulty_level", err))
	}
	if err := val.ValidateCuisineType(req.GetCuisineType()); err != nil {
		violations = append(violations, fieldViolation("cuisine_type", err))
	}
	if err := val.ValidatePositiveInt32(req.GetCalorieCountPerServing(), "calorie_count_per_serving"); err != nil {
		violations = append(violations, fieldViolation("calorie_count_per_serving", err))
	}
	if err := val.ValidatePositiveInt32(req.GetServingsCount(), "servings_count"); err != nil {
		violations = append(violations, fieldViolation("servings_count", err))
	}
	if err := val.ValidateString(req.GetPreparationSteps(), 10, 5000); err != nil {
		violations = append(violations, fieldViolation("preparation_steps", err))
	}
	return violations
}
