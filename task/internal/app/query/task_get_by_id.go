package query

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type GetById struct {
	Id string
}

type GetByIdHandler QueryHandler[GetById, Task]

type getByIdHandler struct {
	repo ReadRepository
}

func (h *getByIdHandler) Handle(ctx context.Context, qry GetById) (Task, error) {
	task, err := h.repo.GetTaskById(ctx, qry.Id)
	if err != nil {
		return Task{}, errors.Wrap(err, fmt.Sprintf("cannot get task by id '%s'", qry.Id))
	}

	return task, nil
}
