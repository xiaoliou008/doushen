package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-demo/service"
	"net/http"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, service.Register(username, password))
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, service.Login(username, password))
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, service.UserInfo(token))
}
