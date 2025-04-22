package main

import (
	"github.com/gin-gonic/gin"
	"ms-mantenedor-clientes/router"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	router.SetupRoutes(r)
	r.Run(":8080") 
}
