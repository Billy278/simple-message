package middleware

import (
	"net/http"
	"simple-message/controller"
	"simple-message/model/web"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	session, _ := controller.Store.Get(ctx.Request, "Auth-Key")
	if len(session.Values) == 0 && ctx.Request.Header.Get("KEY") != "SECRET-KEY" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.ResponseWeb{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		})
		return
	} else if ctx.Request.Header.Get("KEY") == "SECRET-KEY" {
		ctx.Next()
	} else {
		ctx.Next()
	}
}
