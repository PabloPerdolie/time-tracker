package postgres

import (
	"EffectiveMobileTestTask/internal/models"
	"EffectiveMobileTestTask/internal/repository"
	"database/sql"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetByUser(userID int, startDate, endDate string) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0, 32)
	query := `SELECT * FROM tasks WHERE user_id = $1 AND start_time >= $2 AND end_time <= $3 ORDER BY duration DESC`
	rows, err := r.db.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.UserID, &task.Description, &task.StartTime, &task.EndTime, &task.Duration, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *taskRepository) Create(task *models.Task) error {
	query := `
		INSERT INTO tasks 
		(user_id, description, start_time, end_time, duration, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, task.UserID, task.Description, task.StartTime, task.EndTime, task.Duration, task.CreatedAt, task.UpdatedAt)
	return err
}

func (r *taskRepository) Update(task *models.Task) error {
	query := `
		UPDATE tasks SET 
        end_time = $1, duration = $2, updated_at = $3 
		WHERE id = $4`
	_, err := r.db.Exec(query, task.EndTime, task.Duration, task.UpdatedAt, task.ID)
	return err
}
