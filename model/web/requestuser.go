package web

type RequestUser struct {
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
