package model

type CreateRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}

type CreateResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
type PasswordResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UpdateUser struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
