package constant

import "time"

const (
	RefreshTokenExpire = time.Hour * 24 * 7
	AccessTokenExpire  = time.Minute * 5
)
