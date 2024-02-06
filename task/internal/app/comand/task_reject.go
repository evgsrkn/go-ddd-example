package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/task/internal/domain"
	"github.com/pkg/errors"
)

type RejectTask struct {
	Id string
}

type RejectTaskHandler CommandHandler[RejectTask]

type rejectTaskHandler struct {
	repo domain.WriteRepository
}

func (h *rejectTaskHandler) Handle(ctx context.Context, cmd RejectTask) error {
	err := h.repo.UpdateTask(
		ctx,
		cmd.Id,
		func(ctx context.Context, task domain.Task) (domain.Task, error) {
			if err := task.Reject(); err != nil {
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
