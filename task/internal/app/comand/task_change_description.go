package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/task/internal/domain"
	"github.com/pkg/errors"
)

type ChangeDescription struct {
	Id          string
	Description string
}

type ChangeDescriptionHandler CommandHandler[ChangeDescription]

type changeDescriptionHandler struct {
	repo domain.WriteRepository
}

func (h *changeDescriptionHandler) Handle(ctx context.Context, cmd ChangeDescription) error {
	err := h.repo.UpdateTask(
		ctx,
		cmd.Id,
		func(ctx context.Context, task domain.Task) (domain.Task, error) {
			if err := task.Rename(cmd.Description); err != nil {
				return domain.Task{}, err
			}

			return task, nil
		},
	)
	if err != nil {
		return errors.Wrap(err, "cannot rename task")
	}

	return nil
}
