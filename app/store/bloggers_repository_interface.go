package store

import "sumi/app/models"

// BloggersRepository ...
type BloggersRepository interface {
	Create(blogger *models.Blogger) (*models.Blogger, error)
	GetAll() ([]*models.Blogger, error)
	GetByLogin(login string) (*models.Blogger, error)
	Delete() error
	Select(ids []string) error
}
