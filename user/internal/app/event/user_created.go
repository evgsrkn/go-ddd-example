package event

import (
	"time"

	"github.com/evgsrkn/go-ddd-example/user/internal/domain"
)

type UserCreated struct {
	EventBase
	User User
}

type User struct {
	Email        string
	Username     string
	PasswordHash string
	Active       bool
	Role         string
}

func NewUserCreatedEvent(authorId string, user domain.User) UserCreated {
	return UserCreated{
		EventBase{
			AggregateId: user.Id(),
			Author:      authorId,
			Timestamp:   time.Now().UTC(),
		},
		userFromDomain(user),
	}
}

func userFromDomain(user domain.User) User {
	return User{
		Email:        user.Email(),
		Username:     user.Username(),
		PasswordHash: user.PasswordHash(),
		Active:       user.IsActive(),
		Role:         user.Role().String(),
	}
}
