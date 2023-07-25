package gapi

import (
	"fmt"

	db "github.com/PlatosCodes/momsrecipes/db/sqlc"
	"github.com/PlatosCodes/momsrecipes/pb"
	"github.com/PlatosCodes/momsrecipes/token"
	"github.com/PlatosCodes/momsrecipes/util"
)

// Server serves gRPC requests for our moms recipes service.
type Server struct {
	pb.UnimplementedMomsRecipesServer
	config     util.Config
	Store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		Store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
