package domain

import (
	"context"
)

type WriteRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(
		ctx context.Context,
		id string,
		updateFn func(ctx context.Context, user *User) (*User, error),
	) error
}
