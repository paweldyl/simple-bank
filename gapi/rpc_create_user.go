package gapi

import (
	"context"
	"log"

	"github.com/lib/pq"
	db "github.com/paweldyl/simplebank/db/sqlc"
	"github.com/paweldyl/simplebank/pb"
	"github.com/paweldyl/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user already exist: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "faile to create user: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: converUser(user),
	}
	return rsp, nil
}
