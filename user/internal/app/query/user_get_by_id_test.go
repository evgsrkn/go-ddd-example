package query_test

import (
	"context"
	"errors"
	"testing"

	"github.com/evgsrkn/go-ddd-example/user/internal/app/query"
	mock "github.com/evgsrkn/go-ddd-example/user/mocks/user"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserById_success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mock.NewMockReadRepository(ctrl)
	qry := query.NewUserByIdHandler(repo)
	ctx := context.Background()

	expected := newQueryUser(t)

	repo.EXPECT().UserById(ctx, expected.Id).Return(expected, nil)

	actual, err := qry.Handle(ctx, query.UserById{expected.Id})

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestUserById_failure(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mock.NewMockReadRepository(ctrl)
	qry := query.NewUserByIdHandler(repo)
	ctx := context.Background()

	user := newQueryUser(t)

	repo.EXPECT().UserById(ctx, user.Id).Return(nil, errors.New("err"))

	actual, err := qry.Handle(ctx, query.UserById{user.Id})

	assert.Nil(t, actual)
	assert.Error(t, err)
}
