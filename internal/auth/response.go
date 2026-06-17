package auth

import "server/internal/rbac"

type MpAuthResp struct {
	ID     int64     `json:"id"`
	Name   string    `json:"name"`
	Avatar string    `json:"avatar"`
	Role   rbac.Role `json:"role"`
}

type LoginMpResponse struct {
	Token string     `json:"token"`
	Rep   MpAuthResp `json:"rep"`
}
