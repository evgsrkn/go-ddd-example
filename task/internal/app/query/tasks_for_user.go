package query

import (
	"context"

	"github.com/pkg/errors"
)

type TasksForUser struct {
	UserId string
}

type TasksForUserHandler QueryHandler[TasksForUser, []Task]

type tasksForUserHandler struct {
	repo ReadRepository
}

func (h *tasksForUserHandler) Handle(ctx context.Context, qry TasksForUser) ([]Task, error) {
	tasks, err := h.repo.GetTasksForUser(ctx, qry.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get all tasks")
	}

	return tasks, nil
}
