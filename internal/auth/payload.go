package auth

type LoginResponse struct{
	Token string
}
type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}