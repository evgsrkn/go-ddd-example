package domain_test

import (
	"testing"

	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewRole(t *testing.T) {
	t.Parallel()

	for _, role := range domain.UserRoles {
		res, err := domain.RoleFromString(role.String())

		assert.NoError(t, err)
		assert.Equal(t, role, res)
	}
}

func TestNewRole_failure(t *testing.T) {
	t.Parallel()

	_, err := domain.RoleFromString(uuid.NewString())

	assert.Error(t, err)
}

func TestMakeAdmin(t *testing.T) {
	t.Parallel()

	u := createUserExample(t, true)
	err := u.MakeAdmin()

	assert.NoError(t, err)
}

func TestMakeAdmin_failure(t *testing.T) {
	t.Parallel()

	u := createUserExample(t, true)
	_ = u.MakeAdmin()
	err := u.MakeAdmin()

	assert.Error(t, err)
}

func TestIsAdmin(t *testing.T) {
	u := createUserExample(t, true)
	assert.False(t, u.IsAdmin())

	_ = u.MakeAdmin()
	assert.True(t, u.IsAdmin())
}

func TestIsUser(t *testing.T) {
	u := createUserExample(t, true)
	assert.True(t, u.IsUser())

	_ = u.MakeAdmin()
	assert.False(t, u.IsUser())
}
