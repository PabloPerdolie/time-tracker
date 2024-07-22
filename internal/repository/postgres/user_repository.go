package postgres

import (
	"EffectiveMobileTestTask/internal/models"
	"EffectiveMobileTestTask/internal/repository"
	"database/sql"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(filter map[string]interface{}, offset, limit int) ([]*models.User, error) {
	users := make([]*models.User, 0, 32)
	query := `SELECT * FROM users WHERE 1=1`

	for k, v := range filter {
		query = fmt.Sprintf(`%s AND %s = '%v'`, query, k, v)
	}

	query = fmt.Sprintf(`%s OFFSET %d LIMIT %d`, query, offset, limit)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.PassportSeries, &user.PassportNumber, &user.Surname, &user.Name, &user.Patronymic, &user.Address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.PassportSeries, &user.PassportNumber, &user.Surname, &user.Name, &user.Patronymic, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (passport_series, passport_number, surname, name, patronymic, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query, user.PassportSeries, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *userRepository) Update(user *models.User) error {
	query := `
		UPDATE users SET 
		passport_series = $1, 
		passport_number = $2, 
		surname = $3, 
		name = $4, 
		patronymic = $5, 
		address = $6, 
		updated_at = $7 
		WHERE id = $8`
	_, err := r.db.Exec(query,
		user.PassportSeries,
		user.PassportNumber,
		user.Surname,
		user.Name,
		user.Patronymic,
		user.Address,
		user.UpdatedAt,
		user.ID)
	return err
}

func (r *userRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
