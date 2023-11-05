package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
)

type MakeUserAdmin struct {
	Id string
}

type MakeUserAdminHandler CommandHandler[MakeUserAdmin]

type makeUserAdminHandler struct {
	repo domain.WriteRepository
}

func NewMakeUserAdminHandler(repo domain.WriteRepository) MakeUserAdminHandler {
	return makeUserAdminHandler{
		repo: repo,
	}
}

func (h makeUserAdminHandler) Handle(ctx context.Context, cmd MakeUserAdmin) error {
	return h.repo.UpdateUser(
		ctx,
		cmd.Id,
		func(ctx context.Context, user *domain.User) (*domain.User, error) {
			if err := user.MakeAdmin(); err != nil {
				return nil, err
			}

			return user, nil
		},
	)
}
