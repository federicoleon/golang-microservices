package oauth

import (
	"time"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).UTC().Before(time.Now().UTC())
}
