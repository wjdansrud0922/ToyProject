package main

import (
	"github.com/gin-gonic/gin"
	"toyProject/todolist_2/db"
	"toyProject/todolist_2/handlers"
	"toyProject/todolist_2/middlewares"
)

func main() {
	db := db.ConnectDB()
	router := gin.Default()
	router.POST("/sign-up", handlers.Sign_upHandler(db))
	router.POST("/sign-in", handlers.Sign_inHandler(db))

	todoGroup := router.Group("/todo", middlewares.AuthMiddleware())
	{
		todoGroup.GET("", getTodoHandler)
		todoGroup.POST("", createTodoHandler)
		todoGroup.DELETE("/:name", deleteTodoHandler)
		todoGroup.PUT("/:name", updateTodoHandler)
	}
	router.Run(":8080")
}
