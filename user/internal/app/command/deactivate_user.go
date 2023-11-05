package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
)

type DeactivateUser struct {
	Id string
}

type DeactivateUserHandler CommandHandler[DeactivateUser]

type deactivateUserHandler struct {
	repo domain.WriteRepository
}

func NewDeactivateUserHandler(repo domain.WriteRepository) DeactivateUserHandler {
	return deactivateUserHandler{
		repo: repo,
	}
}

func (h deactivateUserHandler) Handle(ctx context.Context, cmd DeactivateUser) error {
	return h.repo.UpdateUser(
		ctx,
		cmd.Id,
		func(ctx context.Context, u *domain.User) (*domain.User, error) {
			if err := u.DeactivateUser(); err != nil {
				return nil, err
			}

			return u, nil
		},
	)
}
