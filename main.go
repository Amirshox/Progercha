package main

import (
	"github.com/gin-gonic/gin"
	"progercha/controllers"
	"progercha/db"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
