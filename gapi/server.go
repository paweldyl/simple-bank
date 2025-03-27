package gapi

import (
	"fmt"

	db "github.com/paweldyl/simplebank/db/sqlc"
	"github.com/paweldyl/simplebank/pb"
	"github.com/paweldyl/simplebank/token"
	"github.com/paweldyl/simplebank/util"
)

// Server service gRPC requests for our banking service
type Server struct {
	// db is the database connection
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
