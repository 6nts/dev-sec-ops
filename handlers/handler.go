package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "¡Hola mundo con Gin!",
	})
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LogIn(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
