package server

import (
	"github.com/gin-gonic/gin"
	"github/binweee/picobin/internal/db"
	"net/http"
)

func Login(c *gin.Context) {
	s, token := db.AuthenticateUser(c.PostForm("username"), c.PostForm("password"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, Success(s))
}
