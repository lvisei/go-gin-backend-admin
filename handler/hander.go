package handler

import (
	"net/http"

	"go-gin-backend-admin/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	g.C.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
