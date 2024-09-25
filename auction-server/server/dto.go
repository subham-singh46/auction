package server

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
