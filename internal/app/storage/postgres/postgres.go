package postgres

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"password-storage-bot/internal/app/models"
	"password-storage-bot/internal/app/storage"
)

type Storage struct {
	db *sql.DB
}

// New creates new postgres storage
func New(databaseUrl string) (*Storage, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return &Storage{db: db}, nil
}

// Set adds service row into db
func (s *Storage) Set(ctx context.Context, service *models.Service) error {
	q := `INSERT INTO services (user_name, service_name, login, password) VALUES ($1, $2, $3, $4)`

	_, err := s.db.ExecContext(ctx, q, service.UserName, service.ServiceName, service.Login, service.Password)

	if err != nil {
		return fmt.Errorf("can't set data: %w ", err)
	}

	return nil
}

// Get gets login and password of service from db
func (s *Storage) Get(ctx context.Context, userName, serviceName string) (*models.Service, error) {
	q := `SELECT login, password FROM services WHERE user_name = $1 AND service_name = $2`

	var (
		login    string
		password string
	)

	err := s.db.QueryRowContext(ctx, q, userName, serviceName).Scan(&login, &password)

	if err == sql.ErrNoRows {
		return nil, storage.ErrNoService
	}

	if err != nil {
		return nil, fmt.Errorf("can't get service %s to user %s: %w", serviceName, userName, err)
	}

	return &models.Service{
		UserName:    userName,
		ServiceName: serviceName,
		Login:       login,
		Password:    password,
	}, nil
}

// Update rewrites login and password for current service
func (s *Storage) Update(ctx context.Context, service *models.Service, newLogin, newPassword string) error {
	q := `UPDATE services SET login = $1, password = $2 WHERE user_name = $3 AND service_name = $4`

	_, err := s.db.ExecContext(ctx, q, newLogin, newPassword, service.UserName, service.ServiceName)

	if err != nil {
		return fmt.Errorf("can't update data: %w ", err)
	}

	return nil
}

// Delete deletes service row
func (s *Storage) Delete(ctx context.Context, userName, serviceName string) error {
	q := `DELETE FROM services WHERE user_name = $1 AND  service_name = $2`

	_, err := s.db.ExecContext(ctx, q, userName, serviceName)

	if err != nil {
		return fmt.Errorf("can't delete data: %w ", err)
	}

	return nil
}

// IsExists checks if service exists in db
func (s *Storage) IsExists(ctx context.Context, service *models.Service) (bool, error) {
	q := `SELECT COUNT(*) FROM services WHERE user_name = $1 AND service_name = $2`

	var count int

	if err := s.db.QueryRowContext(ctx, q, service.UserName, service.ServiceName).Scan(&count); err != nil {
		return false, fmt.Errorf("can't check if service exists: %w", err)
	}

	return count > 0, nil
}
