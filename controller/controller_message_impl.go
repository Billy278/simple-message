package controller

import (
	"fmt"
	"net/http"
	"simple-message/model/web"
	"simple-message/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllerMessageImpl struct {
	Servicemessage  service.ServiceMessage
	Servicelastread service.ServiceLastRead
}

func NewControllerMessageImpl(message service.ServiceMessage, lastread service.ServiceLastRead) ControllerMessage {
	return &ControllerMessageImpl{
		Servicemessage:  message,
		Servicelastread: lastread,
	}
}

func (controller_message *ControllerMessageImpl) SendMessage(c *gin.Context) {
	sendMessage := web.RequestMessage{}
	err := c.ShouldBindJSON(&sendMessage)
	session, _ := Store.Get(c.Request, "Auth-Key")
	sendMessage.Sender = session.Values["username"].(string)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	response := controller_message.Servicemessage.Create(c, sendMessage)
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Pesan Berhasil di kirim",
		Data:   response,
	})
}
func (controller_message *ControllerMessageImpl) SelectPartSender(c *gin.Context) {
	sender := c.Request.URL.Query().Get("sender")
	limit, _ := strconv.Atoi(c.Request.URL.Query().Get("pageSize"))
	offset, _ := strconv.Atoi(c.Request.URL.Query().Get("page"))
	fmt.Println("sender = ", sender)
	fmt.Println("limit = ", limit)
	fmt.Println("offset", offset)
	session, _ := Store.Get(c.Request, "Auth-Key")
	receiver := session.Values["username"].(string)
	Allmessage := controller_message.Servicemessage.SelectPartSender(c, sender, receiver, offset, limit)
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   Allmessage,
	})
}
func (controller_message *ControllerMessageImpl) SelectAllSenderWithLastMessage(c *gin.Context) {
	session, _ := Store.Get(c.Request, "Auth-Key")
	receiver := session.Values["username"].(string)
	fmt.Println(receiver)
	id := controller_message.Servicelastread.FindMaxByid(c)
	fmt.Println(id)
	Allmessage := controller_message.Servicemessage.SelectAllSenderWithLastMessage(c)
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   Allmessage,
	})
}
