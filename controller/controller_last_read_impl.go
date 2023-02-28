package controller

import (
	"fmt"
	"net/http"
	"simple-message/model/domain"
	"simple-message/model/web"
	"simple-message/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllerLastReadImpl struct {
	Servicelastread service.ServiceLastRead
}

func NewControllerLastReadImpl(lastread service.ServiceLastRead) ControllerLastRead {
	return &ControllerLastReadImpl{
		Servicelastread: lastread,
	}

}

func (controller_last *ControllerLastReadImpl) Create(c *gin.Context) {
	var lastread domain.LastRead
	err := c.ShouldBindJSON(&lastread)
	fmt.Println(lastread.Message_id)
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}

	controller_last.Servicelastread.Create(c, lastread.Message_id, lastread.Sender_id)
	id_last := strconv.Itoa(lastread.Message_id)
	c.JSON(http.StatusOK, web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "pesan dengan id = " + id_last + " berhasil di baca",
	})
}

// func (controller_last *ControllerLastReadImpl) FindMaxByid(c *gin.Context) {
// 	controller_last.Servicelastread.FindMaxByid(c)
// }
