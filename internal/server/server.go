package server

import (
	"github.com/gin-gonic/gin"
	"github/binweee/picobin/internal/db"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func Success(data interface{}) Result {
	return Result{
		Code: http.StatusOK,
		Data: data,
	}
}

func Run() {

	db.InitDB()

	route := gin.New()

	r := route.Group("/")

	r.POST("/login", Login)

	route.Run(":8080")
}
