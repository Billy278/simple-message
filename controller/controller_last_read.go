package controller

import "github.com/gin-gonic/gin"

type ControllerLastRead interface {
	Create(c *gin.Context)
	//	FindMaxByid(c *gin.Context)
}
