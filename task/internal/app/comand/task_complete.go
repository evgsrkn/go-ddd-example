package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/task/internal/domain"
	"github.com/pkg/errors"
)

type CompleteTask struct {
	Id string
}

type CompleteTaskHandler CommandHandler[CompleteTask]

type completeTaskHandler struct {
	repo domain.WriteRepository
}

func (h *completeTaskHandler) Handle(ctx context.Context, cmd CompleteTask) error {
	err := h.repo.UpdateTask(
		ctx,
		cmd.Id,
		func(ctx context.Context, task domain.Task) (domain.Task, error) {
			if err := task.Complete(); err != nil {
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
