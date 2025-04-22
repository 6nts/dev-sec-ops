package router

import (
	"github.com/gin-gonic/gin"
	"ms-mantenedor-clientes/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", handlers.IndexHandler)
	r.GET("/hello", handlers.HelloHandler)
}
