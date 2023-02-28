package controller

import "github.com/gin-gonic/gin"

type ControllerUser interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}
