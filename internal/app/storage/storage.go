package storage

import (
	"context"
	"errors"
	"password-storage-bot/internal/app/models"
)

type Storage interface {
	Set(ctx context.Context, service *models.Service) error
	Get(ctx context.Context, userName, serviceName string) (*models.Service, error)
	Update(ctx context.Context, service *models.Service, newLogin, newPassword string) error
	Delete(ctx context.Context, userName, serviceName string) error
	IsExists(ctx context.Context, service *models.Service) (bool, error)
}

var ErrNoService = errors.New("there is no service")
