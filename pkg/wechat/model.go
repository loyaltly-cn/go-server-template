package wechat

type SessionResponse struct {
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid"`
	SessionKey string `json:"session_key"`

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
