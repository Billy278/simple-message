package middleware

import (
	"net/http"
	"simple-message/controller"
	"simple-message/model/web"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	session, _ := controller.Store.Get(ctx.Request, "Auth-Key")
	if len(session.Values) == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.ResponseWeb{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		})
		return
	} else {
		ctx.Next()
	}
}
