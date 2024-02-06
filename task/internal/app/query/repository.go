package query

import "context"

type ReadRepository interface {
	GetTaskById(ctx context.Context, id string) (Task, error)
	GetTasksForUser(ctx context.Context, userId string) ([]Task, error)
	GetAllTasks(ctx context.Context) ([]Task, error)
}
