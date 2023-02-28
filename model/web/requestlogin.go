package web

type RequestLogin struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
