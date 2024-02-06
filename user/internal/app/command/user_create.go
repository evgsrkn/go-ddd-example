package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
	"github.com/pkg/errors"
)

type CreateUser struct {
	Id           string
	Email        string
	Username     string
	PasswordHash string
	Role         string
}

type CreateUserHandler CommandHandler[CreateUser]

type createUserHandler struct {
	repo domain.WriteRepository
}

func NewCreateUserHandler(repo domain.WriteRepository) CreateUserHandler {
	return createUserHandler{
		repo: repo,
	}
}

func (h createUserHandler) Handle(ctx context.Context, cmd CreateUser) (err error) {
	role, err := domain.RoleFromString(cmd.Role)
	if err != nil {
		return errors.Wrap(err, "cannot create user")
	}

	user, err := domain.NewUser(
		cmd.Id,
		cmd.Email,
		cmd.Username,
		cmd.PasswordHash,
		role,
	)
	if err != nil {
		return err
	}

	return h.repo.CreateUser(ctx, user)
}
