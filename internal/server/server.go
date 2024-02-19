package server

import (
	"github.com/gin-gonic/gin"
	"github/binweee/picobin/internal/db"
)

func Run() {

	db.InitDB()

	route := gin.New()

	r := route.Group("/")

	r.POST("/login", Login)

	r.POST("/register", Register)

	route.Run(":8080")
}
