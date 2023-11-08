package query

import "context"

type ReadRepository interface {
	UserById(ctx context.Context, id string) (*User, error)
	AllUsers(ctx context.Context) ([]*User, error)
}
