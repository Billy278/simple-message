package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"simple-message/helper"
	"simple-message/model/web"
	"simple-message/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type ControllerMessageImpl struct {
	Servicemessage  service.ServiceMessage
	Servicelastread service.ServiceLastRead
	Ch              *amqp.Channel
}
type Data struct {
	Date     time.Time `json:"date"`
	Message  string    `json:"message"`
	Receiver string    `json:"receiver"`
}

func NewControllerMessageImpl(message service.ServiceMessage, lastread service.ServiceLastRead, ch *amqp.Channel) ControllerMessage {
	return &ControllerMessageImpl{
		Servicemessage:  message,
		Servicelastread: lastread,
		Ch:              ch,
	}
}

func (controller_message *ControllerMessageImpl) SendMessage(c *gin.Context) {
	sendMessage := web.RequestMessage{}
	err := c.ShouldBindJSON(&sendMessage)
	session, _ := Store.Get(c.Request, "Auth-Key")
	sendMessage.Sender = session.Values["username"].(string)
	sendMessage.Sender_id = session.Values["sender_id"].(int)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	_, err = controller_message.Servicelastread.FindByid(c, sendMessage.Sender_id)
	if err == nil {
		response := controller_message.Servicemessage.Create(c, sendMessage)
		c.JSON(http.StatusOK, web.ResponseWeb{
			Code:   http.StatusOK,
			Status: "Pesan Berhasil di kirim",
			Data:   response,
		})
	} else {
		response := controller_message.Servicemessage.Create(c, sendMessage)
		res := response.Id - 1
		controller_message.Servicelastread.Create(c, res, sendMessage.Sender_id)
		c.JSON(http.StatusOK, web.ResponseWeb{
			Code:   http.StatusOK,
			Status: "Pesan Berhasil di kirim",
			Data:   response,
		})
	}

}

func (controller_message *ControllerMessageImpl) ReceiverPublisher(c *gin.Context) {
	byteData, err := os.ReadFile("./controller/Data.json")
	helper.PanicIfError(err)
	// // 1. Buat koneksi RabbitMQ
	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// helper.PanicIfError(err)
	// defer conn.Close()

	// // 2. Buka channel
	// ch, err := conn.Channel()
	// helper.PanicIfError(err)
	// defer ch.Close()

	// // 3. Deklarasikan exchange
	// err = ch.ExchangeDeclare("my_exchange", "direct", true, false, false, false, nil)
	// helper.PanicIfError(err)

	// //4.declatation queque
	// q, err := ch.QueueDeclare("my_queue", true, false, false, false, nil)
	// helper.PanicIfError(err)

	// // 5. Bind queue ke exchange
	// err = ch.QueueBind(q.Name, "MY_KEY", "my_exchange", false, nil)
	// helper.PanicIfError(err)

	//=======================================
	var data []web.RequestMessage
	json.Unmarshal(byteData, &data)
	for _, v := range data {
		body, _ := json.Marshal(v)
		controller_message.Ch.Publish(
			"my_exchange",   // nama exchange
			"my_routingkey", // routing key
			false,           // mandatory
			false,           // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			},
		)
	}
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Pesan Berhasil Dikirim",
		Data:   data,
	})
}
func (controller_message *ControllerMessageImpl) SendMessageTest(c *gin.Context) {
	// // batas
	H := c.Request.Header.Get("KEY")
	fmt.Println(H)
	sendMessage := web.RequestMessage{}
	err := c.ShouldBindJSON(&sendMessage)
	sendMessage.Sender = "billy"
	sendMessage.Sender_id = 1

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	response := controller_message.Servicemessage.CreateTest(c, sendMessage)
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
	// id := controller_message.Servicelastread.FindMaxByid(c)
	// fmt.Println(id)
	Allmessage := controller_message.Servicemessage.SelectAllSenderWithLastMessage(c, receiver)
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   Allmessage,
	})
}
