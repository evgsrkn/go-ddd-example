package domain

import (
	"context"
)

type WriteRepository interface {
	CreateTask(ctx context.Context, task Task) error
	UpdateTask(
		ctx context.Context,
		id string,
		updateFn func(ctx context.Context, task Task) (Task, error),
	) error
}
