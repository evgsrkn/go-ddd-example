package query

type User struct {
	Id           string
	Email        string
	Username     string
	PasswordHash string
	Active       bool
	Role         string
}
