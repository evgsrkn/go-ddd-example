package query

import (
	"context"
)

type UserById struct {
	Id string
}

type UserByIdHandler QueryHandler[UserById, *User]

type userByIdHandler struct {
	repo ReadRepository
}

func NewUserByIdHandler(repo ReadRepository) UserByIdHandler {
	return userByIdHandler{
		repo: repo,
	}
}

type UserByIdReadModel interface {
	UserById(ctx context.Context, id string) (*User, error)
}

func (h userByIdHandler) Handle(ctx context.Context, qry UserById) (*User, error) {
	return h.repo.UserById(ctx, qry.Id)
}
