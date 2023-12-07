package gapi

import (
	"context"
	"fmt"

	"github.com/PlatosCodes/momsrecipes/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteRecipeByID(ctx context.Context, req *pb.DeleteRecipeByIDRequest) (*emptypb.Empty, error) {
	// Verify the access token
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	err = validateDeleteRecipeRequest(req)
	if err != nil {
		return nil, err
	}

	arg := req.Id

	// Check to make sure recipe exists
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

	// Check if the user ID in the token matches the user id in the recipe
	if authPayload.UserID != recipe.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "you can only delete recipes which you created")
	}

	err = server.Store.DeleteRecipeByID(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteRecipeRequest(req *pb.DeleteRecipeByIDRequest) (err error) {
	if req.Id <= 0 {
		return fmt.Errorf("%d must be positive", req.Id)
	}
	return nil
}
