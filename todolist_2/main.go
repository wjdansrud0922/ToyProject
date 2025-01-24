package main

import (
	"github.com/gin-gonic/gin"
	"toyProject/todolist_2/db"
)

func main() {
	db := db.ConnectDB()
	router := gin.Default()

}
