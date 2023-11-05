package query

import (
	"context"
)

type UserById struct {
	Id string
}

type UserByIdHandler QueryHandler[UserById, *User]

type userByIdHandler struct {
	readModel UserByIdReadModel
}

func NewUserByIdHandler(readModel UserByIdReadModel) UserByIdHandler {
	return userByIdHandler{
		readModel: readModel,
	}
}

type UserByIdReadModel interface {
	UserById(ctx context.Context, id string) (*User, error)
}

func (h userByIdHandler) Handle(ctx context.Context, qry UserById) (*User, error) {
	return h.readModel.UserById(ctx, qry.Id)
}
