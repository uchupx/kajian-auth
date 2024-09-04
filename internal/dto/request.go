package dto

type AuthRequest struct {
	GrantTypePassword
	GrantType    string `json:"grant_type" validate:"required"`
	ClientId     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
}

type GrantTypePassword struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type SignUpRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	ClientKey string `json:"client_key"`
}
