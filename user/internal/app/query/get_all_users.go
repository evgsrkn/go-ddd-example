package query

import (
	"context"
)

type AllUsersHandler QueryHandler[struct{}, []*User]

type allUsersHandler struct {
	readModel AllUsersReadModel
}

func NewAllUsersHandler(readModel AllUsersReadModel) AllUsersHandler {
	return allUsersHandler{
		readModel: readModel,
	}
}

type AllUsersReadModel interface {
	AllUsers(ctx context.Context) ([]*User, error)
}

func (h allUsersHandler) Handle(ctx context.Context, _ struct{}) ([]*User, error) {
	return h.readModel.AllUsers(ctx)
}
