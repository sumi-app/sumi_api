package store

import "sumi/app/models"

// BloggersRepository ...
type BloggersRepository interface {
	Create(blogger *models.Blogger) (*models.Blogger, error)
	GetAll(isSelected bool, isFavorite bool) ([]*models.Blogger, error)
	GetById(id int) (*models.Blogger, error)
	GetByLogin(login string) (*models.Blogger, error)
	Delete() error
	Select(ids []string) error
}
