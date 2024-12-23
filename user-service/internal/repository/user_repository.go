package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"user-service/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uint) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]domain.User, error)
}

type postgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
        INSERT INTO users (first_name, last_name, email, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`

	now := time.Now()
	return r.db.QueryRowContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		now,
		now,
	).Scan(&user.ID)
}

func (r *postgresUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	user := &domain.User{}
	query := `
        SELECT id, first_name, last_name, email, created_at, updated_at
        FROM users
        WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *postgresUserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
        UPDATE users 
        SET first_name = $1, last_name = $2, email = $3, updated_at = $4
        WHERE id = $5`

	result, err := r.db.ExecContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		time.Now(),
		user.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *postgresUserRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *postgresUserRepository) List(ctx context.Context) ([]domain.User, error) {
	query := `
        SELECT id, first_name, last_name, email, created_at, updated_at
        FROM users
        ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
