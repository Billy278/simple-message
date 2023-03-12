package controller

import "github.com/gin-gonic/gin"

type ControllerMessage interface {
	SendMessage(c *gin.Context)
	ReceiverPublisher(c *gin.Context)
	SendMessageTest(c *gin.Context)
	SelectPartSender(c *gin.Context)
	SelectAllSenderWithLastMessage(c *gin.Context)
}
