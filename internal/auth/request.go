package auth

type Provider string

const (
	ProviderWeb Provider = "web"
	ProviderMP  Provider = "mp"
)

type LoginRequest struct {
	Provider Provider `json:"provider" binding:"required"`

	// 小程序
	Code string `json:"code,omitempty"`

	// Web
	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`
}
