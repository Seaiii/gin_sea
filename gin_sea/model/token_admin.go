package model

type TokenAdmin struct {
	Token      string `json:"token"`
	JwtToken   string `json:"jwt_token"`
	CreateTime int64  `json:"create_time"`
	UserId     int    `json:"user_id"`
}
