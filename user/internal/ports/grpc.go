package ports

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/user/internal/app"
	"github.com/evgsrkn/go-ddd-example/user/internal/app/command"
	"github.com/evgsrkn/go-ddd-example/user/internal/app/query"
	"github.com/evgsrkn/go-ddd-example/user/pkg/pb"
)

type GrpcUserServer struct {
	app app.Application
}

func NewGrpcUserServer(app app.Application) GrpcUserServer {
	return GrpcUserServer{
		app: app,
	}
}

func (g GrpcUserServer) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.Empty, error) {
	err := g.app.Commands.CreateUser.Handle(ctx, command.CreateUser{
		Id:           request.Id,
		Email:        request.Email,
		Username:     request.Username,
		PasswordHash: request.PasswordHash,
		Role:         request.Role,
	})
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (g GrpcUserServer) ActivateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.Empty, error) {
	err := g.app.Commands.ActivateUser.Handle(ctx, command.ActivateUser{
		Id: request.Id,
	})
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (g GrpcUserServer) DeactivateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.Empty, error) {
	err := g.app.Commands.DeactivateUser.Handle(ctx, command.DeactivateUser{
		Id: request.Id,
	})
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (g GrpcUserServer) MakeUserAdmin(ctx context.Context, request *pb.UpdateUserRequest) (*pb.Empty, error) {
	err := g.app.Commands.MakeUserAdmin.Handle(ctx, command.MakeUserAdmin{
		Id: request.Id,
	})
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (g GrpcUserServer) GetUserById(ctx context.Context, request *pb.GetUserByIdRequest) (*pb.User, error) {
	user, err := g.app.Queries.GetUserById.Handle(ctx, query.UserById{
		Id: request.Id,
	})
	if err != nil {
		return &pb.User{}, err
	}

	return &pb.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (g GrpcUserServer) GetAllUsers(ctx context.Context, _ *pb.Empty) (*pb.Users, error) {
	users, err := g.app.Queries.GetAllUsers.Handle(ctx, struct{}{})
	if err != nil {
		return &pb.Users{}, err
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	return &pb.Users{Users: pbUsers}, nil
}
