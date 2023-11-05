package domain

import (
	"github.com/pkg/errors"
)

type User struct {
	id           string
	email        string
	username     string
	passwordHash string
	active       bool
	role         Role
}

func NewUser(
	id string,
	email string,
	username string,
	passwordHash string,
	roleStr string,
) (User, error) {
	if id == "" {
		return User{}, errors.New("user id is empty")
	}

	if email == "" {
		return User{}, errors.New("user email is empty")
	}

	if passwordHash == "" {
		return User{}, errors.New("user password is empty")
	}

	role, err := NewRole(roleStr)
	if err != nil {
		return User{}, err
	}

	if role.IsZero() {
		return User{}, errors.New("user role is empty")
	}

	return User{
		id:           id,
		email:        email,
		username:     username,
		passwordHash: passwordHash,
		active:       false,
		role:         role,
	}, nil
}

func UnmarshalUserFromDatabase(
	id string,
	email string,
	username string,
	passwordHash string,
	active bool,
	roleStr string,
) (*User, error) {
	user, err := NewUser(id, email, username, passwordHash, roleStr)
	if err != nil {
		return &User{}, errors.Wrap(err, "cannot unmarshal user from database")
	}

	user.active = active

	return &user, nil
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Role() string {
	return u.role.String()
}

func (u *User) IsActive() bool {
	return u.active
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) ActivateUser() error {
	if u.IsActive() {
		return errors.New("user is already activated")
	}

	u.active = true
	return nil
}

func (u *User) DeactivateUser() error {
	if !u.IsActive() {
		return errors.New("user is already deactivated")
	}

	u.active = false
	return nil
}
