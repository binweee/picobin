package server

import (
	"github.com/gin-gonic/gin"
	"github/binweee/picobin/internal/db"
	"net/http"
)

func Login(c *gin.Context) {
	s := db.AuthenticateUser(c.PostForm("username"), c.PostForm("password"))
	c.JSON(http.StatusOK, s)
}

func Register(c *gin.Context) {
	db.CreateUser(db.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		RoleID:   0,
	})
	c.JSON(http.StatusOK, "success")
}
