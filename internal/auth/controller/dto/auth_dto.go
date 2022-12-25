package dto

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type SignInRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
