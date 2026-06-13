package auth

type UserResponse struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Email  *string `json:"email,omitempty"`
	Avatar string  `json:"avatar"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
