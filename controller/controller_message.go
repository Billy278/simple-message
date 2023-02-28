package controller

import "github.com/gin-gonic/gin"

type ControllerMessage interface {
	SendMessage(c *gin.Context)
	SelectPartSender(c *gin.Context)
	SelectAllSenderWithLastMessage(c *gin.Context)
}
