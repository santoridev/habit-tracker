package main

import (
	"github.com/gin-gonic/gin"
	"github.com/santori/habit-tracker/database"
	"github.com/santori/habit-tracker/handlers"
	_ "modernc.org/sqlite"
)

func main() {
	database.Init()

	r := gin.Default()

	r.POST("/habits", handlers.CreateHabit)
	r.GET("/allhabits", handlers.ReturnAll)
	r.GET("/habitbyid/:id", handlers.ReturnByID)

	r.Run(":8080")
}
