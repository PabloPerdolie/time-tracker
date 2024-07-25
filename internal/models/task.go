package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int          `json:"id"`
	UserID      int          `json:"user_id"`
	Description string       `json:"description"`
	StartTime   time.Time    `json:"start_time"`
	EndTime     sql.NullTime `json:"end_time"`
	Duration    string       `json:"duration"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
