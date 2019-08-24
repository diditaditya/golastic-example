package server

import (
	handlers "app/server/handlers"
	"log"
	
	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()
	r.POST("/documents", handlers.CreateDocuments)
	r.GET("/search", handlers.SearchDocuments)
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}