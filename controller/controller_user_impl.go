package controller

import (
	"net/http"
	"simple-message/helper"
	"simple-message/model/web"
	"simple-message/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type ControllerUserImpl struct {
	Serviceuser service.ServiceUser
}

func NewControllerUserImpl(user service.ServiceUser) ControllerUser {
	return &ControllerUserImpl{
		Serviceuser: user,
	}
}

var key = []byte("Secret-Auth")
var Store = sessions.NewCookieStore(key)

func (controller_user *ControllerUserImpl) Register(c *gin.Context) {
	register_user := web.RequestUser{}
	err := c.ShouldBindJSON(&register_user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	_, err = controller_user.Serviceuser.FindByUser(c, register_user.Username)
	if err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Username Sudah ada",
		})
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(register_user.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	register_user.Password = string(hashPassword)
	controller_user.Serviceuser.Register(c, register_user)
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Registrasi Berhasil",
	})
}
func (controller_user *ControllerUserImpl) Login(c *gin.Context) {
	user := web.RequestLogin{}
	c.ShouldBindJSON(&user)
	users, err := controller_user.Serviceuser.FindByUser(c, user.Username)
	err2 := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password))
	if err != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Username atau password anda salah",
		})
		return
	}
	session, _ := Store.Get(c.Request, "Auth-Key")
	session.Values["username"] = users.Username
	session.Values["sender_id"] = users.Id
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Login berhasil",
	})

}
func (controller_user *ControllerUserImpl) Logout(c *gin.Context) {
	session, _ := Store.Get(c.Request, "Auth-Key")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "anda Berhasil Logout",
	})

}
