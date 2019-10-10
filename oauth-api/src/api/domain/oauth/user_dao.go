package oauth

import (
	"github.com/federicoleon/golang-microservices/src/api/utils/errors"
)

const (
	queryQetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"fede": {Id: 123, Username: "fede"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}
	return user, nil
}

