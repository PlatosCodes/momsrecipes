package gapi

import (
	"context"

	db "github.com/PlatosCodes/momsrecipes/db/sqlc"
	"github.com/PlatosCodes/momsrecipes/pb"
	"github.com/PlatosCodes/momsrecipes/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username: req.GetUsername(),
		Password: hashedPassword,
		Email:    req.GetEmail(),
	}

	user, err := server.Store.RegisterTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertRegisterUser(user),
	}
	return rsp, nil
}
