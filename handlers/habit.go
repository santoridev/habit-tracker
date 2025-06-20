package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/santori/habit-tracker/database"
	"github.com/santori/habit-tracker/models"
	_ "modernc.org/sqlite"
)

func CreateHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	habit.CreatedAt = time.Now()

	_, err := database.DB.Exec(
		`INSERT INTO habits (name, description , goal_days, created_at) VALUES (?,?,?,?)`, habit.Name, habit.Description, habit.GoalDays, habit.CreatedAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	c.JSON(http.StatusOK, habit)
}

func ReturnAll(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT id, name, description, goal_days, created_at FROM habits`)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no rows available"})
		return
	}
	defer rows.Close()
	var habits []models.Habit
	for rows.Next() {
		var h models.Habit
		err := rows.Scan(&h.ID, &h.Name, &h.Description, &h.GoalDays, &h.CreatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing a row"})
			return
		}
		habits = append(habits, h)
	}
	c.JSON(http.StatusOK, habits)
}

func ReturnByID(c *gin.Context) {
	id := c.Param("id")

	var habbit models.Habit

	err := database.DB.QueryRow(`SELECT id, name, description, goal_days, created_at FROM habits WHERE id = ?`, id).Scan(&habbit.ID, &habbit.Name, &habbit.Description, &habbit.GoalDays, &habbit.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "this row doesn't exist"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}
	c.JSON(http.StatusOK, habbit)

}
