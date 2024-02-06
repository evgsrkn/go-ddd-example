package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
)

type ActivateUser struct {
	Id string
}

type ActivateUserHandler CommandHandler[ActivateUser]

type activateUserHandler struct {
	repo domain.WriteRepository
}

func NewActivateUserHandler(repo domain.WriteRepository) ActivateUserHandler {
	return activateUserHandler{
		repo: repo,
	}
}

func (h activateUserHandler) Handle(ctx context.Context, cmd ActivateUser) error {
	return h.repo.UpdateUser(
		ctx,
		cmd.Id,
		func(ctx context.Context, user *domain.User) (*domain.User, error) {
			if err := user.ActivateUser(); err != nil {
				return nil, err
			}

			return user, nil
		},
	)
}
