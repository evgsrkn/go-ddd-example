package query

import (
	"context"

	"github.com/pkg/errors"
)

type GetAll struct{}

type GetAllHandler QueryHandler[GetAll, []Task]

type getAllHandler struct {
	repo ReadRepository
}

func (h *getAllHandler) Handle(ctx context.Context, qry GetAll) ([]Task, error) {
	tasks, err := h.repo.GetAllTasks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get all tasks")
	}

	return tasks, nil
}
