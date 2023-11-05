package service

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/user/internal/adapters/postgresql"
	"github.com/evgsrkn/go-ddd-example/user/internal/app"
	"github.com/evgsrkn/go-ddd-example/user/internal/app/command"
	"github.com/evgsrkn/go-ddd-example/user/internal/app/query"
	"github.com/evgsrkn/go-ddd-example/user/internal/config"
	"github.com/evgsrkn/go-ddd-example/user/internal/ports"
	"github.com/evgsrkn/go-ddd-example/user/internal/server"

	"github.com/evgsrkn/go-ddd-example/gateway/pkg/logger"
	"github.com/evgsrkn/go-ddd-example/gateway/pkg/rpc"
)

func NewApplication() {
	ctx := context.Background()
	cfg := config.New()
	log := logger.New(cfg.AppEnv.String())

	conn := postgresql.NewConnectionPool(ctx, cfg, log)
	db := postgresql.NewPostgresqlRepository(conn)

	app := app.Application{
		Queries: app.Queries{
			GetUserById: query.NewUserByIdHandler(db),
			GetAllUsers: query.NewAllUsersHandler(db),
		},
		Commands: app.Commands{
			ActivateUser:   command.NewActivateUserHandler(db),
			DeactivateUser: command.NewDeactivateUserHandler(db),
			MakeUserAdmin:  command.NewMakeUserAdminHandler(db),
			CreateUser:     command.NewCreateUserHandler(db),
		},
	}

	usrSrv := ports.NewGrpcUserServer(app)
	srv := rpc.New(
		rpc.WithLoggerZap(log),
	)

	server.ServeGrpc(log, srv, &usrSrv)
}
