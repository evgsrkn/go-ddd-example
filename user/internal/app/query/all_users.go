package query

import (
	"context"
)

type AllUsers struct{}

type AllUsersHandler QueryHandler[AllUsers, []*User]

type allUsersHandler struct {
	repo ReadRepository
}

func NewAllUsersHandler(repo ReadRepository) AllUsersHandler {
	return allUsersHandler{
		repo: repo,
	}
}

func (h allUsersHandler) Handle(ctx context.Context, _ AllUsers) ([]*User, error) {
	return h.repo.AllUsers(ctx)
}
