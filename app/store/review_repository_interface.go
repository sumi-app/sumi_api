package store

import "sumi/app/models"

// ReviewsRepository ...
type ReviewsRepository interface {
	Create(review *models.Review) (*models.Review, error)
	GetAll() ([]*models.Review, error)
	GetByPressTourId(id int) ([]*models.Review, error)
	Delete() error
}
