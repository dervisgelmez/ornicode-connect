package response

import "time"

type OAuthTokenResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}
