package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-message/app"
	"simple-message/controller"
	"simple-message/middleware"
	"simple-message/repository"
	"simple-message/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	conn, ch := app.NewRabbitmq()
	defer conn.Close()
	defer ch.Close()
	DB := app.NewDB()
	Validate := validator.New()
	RepositoryUser := repository.NewRepositoryUserImpl()
	RepositoryMessage := repository.NewRepositoryMessageImpl()
	RepositoryLast := repository.NewRepositoryLastReadImpl()

	ServiceUser := service.NewServiceUserImpl(DB, Validate, RepositoryUser)
	ServiceMessage := service.NewServiceMessageImpl(DB, Validate, RepositoryMessage)
	ServiceLast := service.NewServiceLastReadImpl(DB, RepositoryLast)

	ControllerUser := controller.NewControllerUserImpl(ServiceUser)
	ControllerMessage := controller.NewControllerMessageImpl(ServiceMessage, ServiceLast, ch)
	ControllerLast := controller.NewControllerLastReadImpl(ServiceLast)

	if os.Getenv("MODE") == "API" {
		fmt.Println("api")
		// jalankan server
		router.Use(gin.Recovery())
		router.POST("/register", ControllerUser.Register)
		router.POST("/login", ControllerUser.Login)
		router.POST("/messageTest", ControllerMessage.SendMessageTest)
		authRouter := router.Group("/users")
		authRouter.Use(middleware.AuthMiddleware)
		authRouter.GET("/message", ControllerMessage.SelectAllSenderWithLastMessage)

		authRouter.GET("/publisher/sender", ControllerMessage.ReceiverPublisher)
		authRouter.POST("/messageTest", ControllerMessage.SendMessageTest)
		authRouter.POST("/message", ControllerMessage.SendMessage)
		authRouter.GET("/read", ControllerMessage.SelectPartSender)
		authRouter.POST("/setread", ControllerLast.Create)
		authRouter.GET("/logout", ControllerUser.Logout)

		server := http.Server{
			Addr:    "localhost:9001",
			Handler: router,
		}

		server.ListenAndServe()
	}
	if os.Getenv("MODE") == "SUBSCRIBER" {
		fmt.Println("Subcriber")

		// // jalankan subscriber
		// // 1. Buat koneksi RabbitMQ
		// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		// if err != nil {
		// 	log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		// }
		// defer conn.Close()

		// // 2. Buka channel
		// ch, err := conn.Channel()
		// if err != nil {
		// 	log.Fatalf("Failed to open a channel: %v", err)
		// }
		// defer ch.Close()

		// 7. Konsumsi pesan dari queue
		msgs, err := ch.Consume(
			//q.Name, // nama queue
			"my_queue",
			"",    // consumer
			true,  // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // argument
		)
		if err != nil {
			log.Fatalf("Failed to register a consumer: %v", err)
		}
		//var data []Mahasiswa
		// 8. Looping untuk mengambil pesan
		for msg := range msgs {
			fmt.Printf("Received a message: %s\n", msg.Body)
			reader := bytes.NewReader(msg.Body)
			//test perbedaaan
			//http.Post("http://localhost:9001/messageTest", "application/json", reader)

			req, err := http.NewRequest("POST", "http://localhost:9001/users/messageTest", reader)
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Add("KEY", "SECRET-KEY")
			http.DefaultClient.Do(req)

		}

	}
	fmt.Println("test")

}
