package store

import "sumi/app/models"

// BloggersRepository ...
type BloggersRepository interface {
	Create(blogger *models.Blogger) (*models.Blogger, error)
	GetAll() ([]*models.Blogger, error)
	Delete() error
}
