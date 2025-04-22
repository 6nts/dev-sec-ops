package router

import (
	"ms-mantenedor-clientes/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", handlers.IndexHandler)
	r.GET("/login", handlers.LogIn)
	r.GET("/hello", handlers.HelloHandler)
}
