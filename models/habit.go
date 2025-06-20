package models

import "time"

type Habit struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	GoalDays    int       `json:"goal"`
	CreatedAt   time.Time `json:"createdat"`
}
