package query_test

import (
	"context"
	"errors"
	"testing"

	"github.com/evgsrkn/go-ddd-example/user/internal/app/query"
	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
	mock "github.com/evgsrkn/go-ddd-example/user/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAllUsers_success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mock.NewMockReadRepository(ctrl)
	qry := query.NewAllUsersHandler(repo)
	ctx := context.Background()

	expected := []*query.User{
		newQueryUser(t),
		newQueryUser(t),
	}

	repo.EXPECT().AllUsers(ctx).Return(expected, nil)

	actual, err := qry.Handle(ctx, query.AllUsers{})

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestAllUsers_failure(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mock.NewMockReadRepository(ctrl)
	qry := query.NewAllUsersHandler(repo)
	ctx := context.Background()

	repo.EXPECT().AllUsers(ctx).Return(nil, errors.New("err"))

	actual, err := qry.Handle(ctx, query.AllUsers{})

	assert.Nil(t, actual)
	assert.Error(t, err)
}

func newQueryUser(t *testing.T) *query.User {
	t.Helper()

	return &query.User{
		Id:           uuid.NewString(),
		Email:        "example@mail.com",
		Username:     "username",
		PasswordHash: uuid.NewString(),
		Active:       true,
		Role:         domain.UserRole.String(),
	}
}
