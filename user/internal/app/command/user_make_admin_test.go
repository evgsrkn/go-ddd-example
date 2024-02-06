package command_test

import (
	"context"
	"errors"
	"testing"

	"github.com/evgsrkn/go-ddd-example/user/internal/app/command"
	mock "github.com/evgsrkn/go-ddd-example/user/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMakeUserAdmin(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mock.NewMockWriteRepository(ctrl)
	cmd := command.NewMakeUserAdminHandler(repo)
	ctx := context.Background()

	userId := uuid.NewString()

	repo.EXPECT().UpdateUser(
		ctx,
		userId,
		gomock.Any(),
	).Return(nil)

	err := cmd.Handle(ctx, command.MakeUserAdmin{Id: userId})
	assert.NoError(t, err)
}

func TestMakeUserAdmin_failure(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mock.NewMockWriteRepository(ctrl)
	cmd := command.NewMakeUserAdminHandler(repo)
	ctx := context.Background()

	userId := uuid.NewString()

	repo.EXPECT().UpdateUser(
		ctx,
		userId,
		gomock.Any(),
	).Return(errors.New("err"))

	err := cmd.Handle(ctx, command.MakeUserAdmin{Id: userId})
	assert.Error(t, err)
}
