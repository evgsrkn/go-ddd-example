package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/task/internal/domain"
	"github.com/pkg/errors"
)

type RenameTask struct {
	Id   string
	Name string
}

type RenameTaskHandler CommandHandler[CreateTask]

type renameTaskHandler struct {
	repo domain.WriteRepository
}

func (h *renameTaskHandler) Handle(ctx context.Context, cmd RenameTask) error {
	err := h.repo.UpdateTask(
		ctx,
		cmd.Id,
		func(ctx context.Context, task domain.Task) (domain.Task, error) {
			if err := task.Rename(cmd.Name); err != nil {
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
