package store

// Store ...
type Store interface {
	Blogger() BloggersRepository
	Review() ReviewsRepository
	PressTour() PressTourRepository
}
