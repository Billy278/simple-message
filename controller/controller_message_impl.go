package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-message/model/web"
	"simple-message/service"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

type ControllerMessageImpl struct {
	Servicemessage  service.ServiceMessage
	Servicelastread service.ServiceLastRead
}
type Data struct {
	Date     time.Time `json:"date"`
	Message  string    `json:"message"`
	Receiver string    `json:"receiver"`
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
	sendMessage.Sender_id = session.Values["sender_id"].(int)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	// _, err = controller_message.Servicelastread.FindByid(c, sendMessage.Sender_id)
	// if err == nil {
	// 	response := controller_message.Servicemessage.Create(c, sendMessage)
	// 	c.JSON(http.StatusOK, web.ResponseWeb{
	// 		Code:   http.StatusOK,
	// 		Status: "Pesan Berhasil di kirim",
	// 		Data:   response,
	// 	})
	// } else {
	response := controller_message.Servicemessage.Create(c, sendMessage)
	res := response.Id - 1
	controller_message.Servicelastread.Create(c, res, sendMessage.Sender_id)
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Pesan Berhasil di kirim",
		Data:   response,
	})
	//}

}

func (controller_message *ControllerMessageImpl) ReceiverPublisher(c *gin.Context) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Read data from JSON file
	dataBytes, err := os.ReadFile("./controller/Data.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(dataBytes))

	// Parse JSON data into a Data object
	var data []Data
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		log.Fatal(err)
	}
	// date, _ := time.Parse(time.Layout, "2023-03-14T00:00:00Z")
	// data := []Data{
	// 	Data{
	// 		Date:     date,
	// 		Message:  "hallo test 1",
	// 		Receiver: "bima",
	// 	},
	// }
	//fmt.Println(data)
	//Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)
	//Subscribe
	if _, err := nc.Subscribe("sample_subject", func(m *nats.Msg) {

		fmt.Println("message received : ", string(m.Data))
		requestBody := strings.NewReader(`{"date" : "2023-03-14T00:00:00Z","message" : "Gadget","receiver":"bima"}`)

		resp, err := http.Post("http://localhost:9001/users/messageTest", "application/json", requestBody)

		// c.Request.Form.Add("date", "2023-03-14T00:00:00Z")
		// c.Request.Form.Add("message", "Test")
		// c.Request.Form.Add("receiver", "bima")

		//http.Redirect(c.Writer, c.Request, "/users/messageTest", http.StatusSeeOther)
		if err != nil {

			fmt.Println(err)
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.ResponseWeb{
				Code:   http.StatusOK,
				Status: "Tes",
			})
		}
		fmt.Println(resp.Status)
		fmt.Println("success")
		// defer resp Close()

		var dt Data
		err = json.Unmarshal(m.Data, &dt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dt)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}
	for _, v := range data {
		dataJSON, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		if err := nc.Publish("sample_subject", dataJSON); err != nil {
			log.Fatal(err)
			log.Println(data)
		}
	}
	wg.Wait()
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Tes",
	})
	//http.NewRequest(http.MethodPost, "http://localhost:9001/users/messageTest", requestBody)
	//r, err := http.Post("http://localhost:9001/users/messageTest", "application/json", requestBody)
	//http.Redirect(c.Writer, r, "/users/messageTest", http.StatusSeeOther)
}
func (controller_message *ControllerMessageImpl) SendMessageTest(c *gin.Context) {
	fmt.Println("test=========================================")
	// // batas
	// sendMessage := web.RequestMessage{}
	// err := c.ShouldBindJSON(&sendMessage)
	// session, _ := Store.Get(c.Request, "Auth-Key")
	// sendMessage.Sender = session.Values["username"].(string)
	// sendMessage.Sender_id = session.Values["sender_id"].(int)

	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
	// 		Code:   http.StatusBadRequest,
	// 		Status: "Bad Request",
	// 		Data:   err,
	// 	})
	// 	return
	// }
	// response := controller_message.Servicemessage.CreateTest(c, sendMessage)
	// c.JSON(http.StatusOK, web.ResponseWeb{
	// 	Code:   http.StatusOK,
	// 	Status: "Pesan Berhasil di kirim",
	// 	Data:   response,
	// })

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
