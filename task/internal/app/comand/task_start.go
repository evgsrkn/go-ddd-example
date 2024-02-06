package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/task/internal/domain"
	"github.com/pkg/errors"
)

type StartTask struct {
	Id string
}

type StartTaskHandler CommandHandler[StartTask]

type startTaskHandler struct {
	repo domain.WriteRepository
}

func (h *startTaskHandler) Handle(ctx context.Context, cmd StartTask) error {
	err := h.repo.UpdateTask(
		ctx,
		cmd.Id,
		func(ctx context.Context, task domain.Task) (domain.Task, error) {
			if err := task.Start(); err != nil {
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
