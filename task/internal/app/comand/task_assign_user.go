package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/task/internal/domain"
	userPb "github.com/evgsrkn/go-ddd-example/user/pkg/pb"
	"github.com/pkg/errors"
)

type AssignUser struct {
	Id     string
	UserId string
}

type AssignUserHandler CommandHandler[AssignUser]

type assignUserHandler struct {
	repo    domain.WriteRepository
	userSvc userPb.UserServiceClient
}

func (h *assignUserHandler) Handle(ctx context.Context, cmd AssignUser) error {
	_, err := h.userSvc.GetUserById(ctx, &userPb.GetUserByIdRequest{Id: cmd.UserId})
	if err != nil {
		return errors.Wrap(err, "cannot assign user")
	}

	err = h.repo.UpdateTask(
		ctx,
		cmd.Id,
		func(ctx context.Context, task domain.Task) (domain.Task, error) {
			if err := task.Assign(cmd.UserId); err != nil {
				return domain.Task{}, err
			}

			return task, nil
		},
	)
	if err != nil {
		return errors.Wrap(err, "cannot assign user")
	}

	return nil
}
