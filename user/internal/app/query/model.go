package query

import "github.com/evgsrkn/go-ddd-example/user/internal/domain"

type User struct {
	Id           string
	Email        string
	Username     string
	PasswordHash string
	Active       bool
	Role         domain.Role
}
