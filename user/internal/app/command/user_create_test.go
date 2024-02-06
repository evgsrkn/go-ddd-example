package command_test

import (
	"context"
	"errors"
	"testing"

	"github.com/evgsrkn/go-ddd-example/user/internal/app/command"
	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
	mock "github.com/evgsrkn/go-ddd-example/user/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		setup         func(repo *mock.MockWriteRepository, cmd *command.CreateUser)
		errorExpected bool
	}{
		"TestCreateUser success": {
			func(repo *mock.MockWriteRepository, cmd *command.CreateUser) {
				repo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			false,
		},
		"TestCreateUser failure": {
			func(repo *mock.MockWriteRepository, cmd *command.CreateUser) {
				repo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(errors.New("err"))
			},
			true,
		},
		"TestCreateUser with empty username": {
			func(repo *mock.MockWriteRepository, cmd *command.CreateUser) {
				cmd.Username = ""
			},
			true,
		},
		"TestCreateUser with invalid role": {
			func(repo *mock.MockWriteRepository, cmd *command.CreateUser) {
				cmd.Role = uuid.NewString()
			},
			true,
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			repo := mock.NewMockWriteRepository(ctrl)
			handler := command.NewCreateUserHandler(repo)
			ctx := context.Background()

			cmd := createValidCommand(t)
			test.setup(repo, cmd)

			err := handler.Handle(ctx, *cmd)

			if test.errorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func createValidCommand(t *testing.T) *command.CreateUser {
	t.Helper()

	return &command.CreateUser{
		Id:           uuid.NewString(),
		Email:        "test@email.com",
		Username:     "username",
		PasswordHash: uuid.NewString(),
		Role:         domain.UserRole.String(),
	}
}
