package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	AppID  string
	Secret string
}

func NewClient(appID, secret string) *Client {
	return &Client{
		AppID:  appID,
		Secret: secret,
	}
}

func (c *Client) Code2Session(code string) (*SessionResponse, error) {

	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		c.AppID,
		c.Secret,
		code,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SessionResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf(
			"wechat error: %d %s",
			result.ErrCode,
			result.ErrMsg,
		)
	}

	return &result, nil
}
