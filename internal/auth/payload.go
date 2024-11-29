package auth

type LoginResponse struct{
	Token string
}
type LoginRequest struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}