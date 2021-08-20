package sqlstore

import (
	"database/sql"
	"sumi/app/store"

	_ "github.com/lib/pq" // ...
)

type Store struct {
	db                *sql.DB
	bloggerRepository store.BloggersRepository
	reviewsRepository store.ReviewsRepository
	pressTourRepository store.PressTourRepository
}

//New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//Blogger ...
func (s *Store) Blogger() store.BloggersRepository {
	if s.bloggerRepository == nil {
		s.bloggerRepository = &BloggersRepository{
			store: s,
		}
	}
	return s.bloggerRepository
}

func (s *Store) Review() store.ReviewsRepository {
	if s.reviewsRepository == nil{
		s.reviewsRepository = &ReviewsRepository{
		store: s,
	}
	}
	return s.reviewsRepository
}

func (s *Store) PressTour() store.PressTourRepository {
	if s.pressTourRepository == nil{
		s.pressTourRepository = &PressTourRepository{
			store: s,
		}
	}
	return s.pressTourRepository
}