package domain_test

import (
	"testing"

	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	Id           string
	Email        string
	Username     string
	PasswordHash string
	Active       bool
	Role         domain.Role
}

func TestNewUser(t *testing.T) {
	t.Parallel()

	u := createTestUser(t)
	usr, err := domain.NewUser(u.Id, u.Email, u.Username, u.PasswordHash, u.Role)
	assert.NoError(t, err)

	assert.Equal(t, u.Id, usr.Id())
	assert.Equal(t, u.Email, usr.Email())
	assert.Equal(t, u.Username, usr.Username())
	assert.Equal(t, u.PasswordHash, usr.PasswordHash())
	assert.False(t, usr.IsActive())
	assert.Equal(t, u.Role, usr.Role())
}

func TestNewUser_failure(t *testing.T) {
	t.Parallel()

	u := createTestUser(t)

	_, err := domain.NewUser("", u.Email, u.Username, u.PasswordHash, u.Role)
	assert.Error(t, err)

	_, err = domain.NewUser(u.Id, "", u.Username, u.PasswordHash, u.Role)
	assert.Error(t, err)

	_, err = domain.NewUser(u.Id, u.Email, "", u.PasswordHash, u.Role)
	assert.Error(t, err)

	_, err = domain.NewUser(u.Id, u.Email, u.Username, "", u.Role)
	assert.Error(t, err)

	_, err = domain.NewUser(u.Id, u.Email, u.Username, u.PasswordHash, domain.Role{})
	assert.Error(t, err)
}

func TestActivateUser(t *testing.T) {
	user := createUserExample(t, false)
	err := user.ActivateUser()

	assert.NoError(t, err)
	assert.True(t, user.IsActive())
}

func TestActivateUser_failure(t *testing.T) {
	user := createUserExample(t, true)
	err := user.ActivateUser()

	assert.Error(t, err)
	assert.True(t, user.IsActive())
}

func TestDeactivateUser(t *testing.T) {
	user := createUserExample(t, true)
	err := user.DeactivateUser()

	assert.NoError(t, err)
	assert.False(t, user.IsActive())
}

func TestDeactivateUser_failure(t *testing.T) {
	user := createUserExample(t, false)
	err := user.DeactivateUser()

	assert.Error(t, err)
	assert.False(t, user.IsActive())
}

func TestUnmarshalUserFromDatabase(t *testing.T) {
	t.Parallel()

	u := createTestUser(t)

	usr, err := domain.UnmarshalUserFromDatabase(
		u.Id,
		u.Email,
		u.Username,
		u.PasswordHash,
		u.Active,
		u.Role,
	)

	assert.NoError(t, err)
	assert.Equal(t, u.Id, usr.Id())
	assert.Equal(t, u.Email, usr.Email())
	assert.Equal(t, u.Username, usr.Username())
	assert.Equal(t, u.PasswordHash, usr.PasswordHash())
	assert.Equal(t, u.Active, usr.IsActive())
	assert.Equal(t, u.Role, usr.Role())
}

func TestUnmarshalUserFromDatabase_failure(t *testing.T) {
	t.Parallel()

	u := createTestUser(t)

	_, err := domain.UnmarshalUserFromDatabase(
		u.Id,
		"",
		u.Username,
		u.PasswordHash,
		u.Active,
		u.Role,
	)

	assert.Error(t, err)
}

func createTestUser(t *testing.T) *TestUser {
	t.Helper()

	return &TestUser{
		Id:           uuid.NewString(),
		Email:        "test@mail.com",
		Username:     "username",
		Active:       true,
		PasswordHash: uuid.NewString(),
		Role:         domain.UserRole,
	}
}

func createUserExample(t *testing.T, isActive bool) *domain.User {
	t.Helper()

	usr, err := domain.NewUser(
		uuid.NewString(),
		"test@mail.com",
		"username",
		uuid.NewString(),
		domain.UserRole,
	)
	assert.NoError(t, err)

	if isActive {
		err = usr.ActivateUser()
		assert.NoError(t, err)
	}

	return usr
}
