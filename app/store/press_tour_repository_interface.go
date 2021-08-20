package store

import "sumi/app/models"

// PressTourRepository ...
type PressTourRepository interface {
	Create(blogger *models.PressTour) (*models.PressTour, error)
	GetAll() ([]*models.PressTour, error)
	Delete() error
}
