package app

import (
	"github.com/federicoleon/golang-microservices/src/api/controllers/polo"
	"github.com/federicoleon/golang-microservices/oauth-api/src/api/controllers/oauth"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
