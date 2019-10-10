package oauth

import (
	"github.com/federicoleon/golang-microservices/src/api/utils/errors"
	"fmt"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() (errors.ApiError) {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.ApiError) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.NewNotFoundError("no access token found with given parameters")
	}
	return token, nil
}
