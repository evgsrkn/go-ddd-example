package domain

import "github.com/pkg/errors"

var (
	AdminRole = Role{"admin"}
	UserRole  = Role{"user"}
)

var UserRoles = []Role{
	AdminRole,
	UserRole,
}

type Role struct {
	a string
}

func RoleFromString(role string) (Role, error) {
	for _, r := range UserRoles {
		if role == r.String() {
			return r, nil
		}
	}

	return Role{}, errors.Errorf("unknown '%s' role", role)
}

func (r Role) String() string {
	return r.a
}

func (r Role) IsZero() bool {
	return r == Role{}
}

func (u User) IsAdmin() bool {
	return u.role == AdminRole
}

func (u User) IsUser() bool {
	return u.role == UserRole
}

func (u *User) MakeAdmin() error {
	if u.IsAdmin() {
		return errors.New("user is already an admin")
	}

	u.role = AdminRole
	return nil
}
