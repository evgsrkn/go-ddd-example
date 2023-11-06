package domain

import (
	"context"
)

type ViewRepository interface {
	GetUserById(ctx context.Context, id string) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}

type WriteRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(
		ctx context.Context,
		id string,
		updateFn func(ctx context.Context, user *User) (*User, error),
	) error
}
