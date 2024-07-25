package postgres

import (
	"EffectiveMobileTestTask/internal/models"
	"EffectiveMobileTestTask/internal/repository"
	"database/sql"
	"strconv"
	"time"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) GetByUser(userID int, startDate, endDate string) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0, 32)
	query := `SELECT * FROM tasks WHERE user_id = $1 AND start_time >= $2 AND end_time <= $3 ORDER BY duration DESC`
	rows, err := t.db.Query(query, userID, startDate, endDate)
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

func (t *taskRepository) Create(userID int, desc string) (*models.Task, error) {
	task := models.Task{
		UserID:      userID,
		Description: desc,
		StartTime:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	query := `
		INSERT INTO tasks 
		(user_id, description, start_time, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := t.db.QueryRow(query,
		task.UserID, task.Description, task.StartTime, task.CreatedAt, task.UpdatedAt).
		Scan(&task.ID)
	if err != nil {
		return nil, err
	}
	return &task, err
}

func (t *taskRepository) Update(id int) (*models.Task, error) {
	task, err := t.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	duration := int(time.Since(task.StartTime).Seconds())
	task.Duration = strconv.Itoa(duration)
	task.EndTime = sql.NullTime{Time: time.Now(), Valid: true}
	task.UpdatedAt = time.Now()
	query := `
		UPDATE tasks SET 
        end_time = $1, duration = $2, updated_at = $3 
		WHERE id = $4`
	_, err = t.db.Exec(query, task.EndTime, task.Duration, task.UpdatedAt, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskRepository) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	query := `SELECT * FROM tasks WHERE id = $1`
	err := t.db.QueryRow(query, id).Scan(&task.ID, &task.UserID, &task.Description, &task.StartTime, &task.EndTime, &task.Duration, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
