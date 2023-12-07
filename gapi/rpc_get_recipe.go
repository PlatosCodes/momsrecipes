package gapi

import (
	"context"
	"fmt"

	"github.com/PlatosCodes/momsrecipes/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetRecipeByID(ctx context.Context, req *pb.GetRecipeByIDRequest) (*pb.GetRecipeByIDResponse, error) {
	err := validateGetRecipeRequest(req)
	if err != nil {
		return nil, err
	}

	arg := req.Id

	recipe, err := server.Store.GetRecipeByID(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "case_not_found":
				return nil, status.Errorf(codes.NotFound, "recipe with that id does not exist: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	return &pb.GetRecipeByIDResponse{
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

func validateGetRecipeRequest(req *pb.GetRecipeByIDRequest) (err error) {
	if req.Id <= 0 {
		return fmt.Errorf("%d must be positive", req.Id)
	}
	return nil
}
