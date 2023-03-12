package main

import (
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
)

func main() {
	if os.Getenv("MODE") == "API" {
		// jalankan server
	} else {
		// jalankan subscriber
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	DB := app.NewDB()
	Validate := validator.New()
	RepositoryUser := repository.NewRepositoryUserImpl()
	RepositoryMessage := repository.NewRepositoryMessageImpl()
	RepositoryLast := repository.NewRepositoryLastReadImpl()

	ServiceUser := service.NewServiceUserImpl(DB, Validate, RepositoryUser)
	ServiceMessage := service.NewServiceMessageImpl(DB, Validate, RepositoryMessage)
	ServiceLast := service.NewServiceLastReadImpl(DB, RepositoryLast)

	ControllerUser := controller.NewControllerUserImpl(ServiceUser)
	ControllerMessage := controller.NewControllerMessageImpl(ServiceMessage, ServiceLast)
	ControllerLast := controller.NewControllerLastReadImpl(ServiceLast)

	router.Use(gin.Recovery())
	router.POST("/register", ControllerUser.Register)
	router.POST("/login", ControllerUser.Login)
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
