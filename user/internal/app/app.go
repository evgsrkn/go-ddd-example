package app

import (
	"github.com/evgsrkn/go-ddd-example/user/internal/app/command"
	"github.com/evgsrkn/go-ddd-example/user/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ActivateUser   command.ActivateUserHandler
	DeactivateUser command.DeactivateUserHandler
	MakeUserAdmin  command.MakeUserAdminHandler
	CreateUser     command.CreateUserHandler
}

type Queries struct {
	GetUserById query.UserByIdHandler
	GetAllUsers query.AllUsersHandler
}
