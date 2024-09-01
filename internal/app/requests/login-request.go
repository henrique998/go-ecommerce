package requests

type LoginRequest struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}
