package sqlstore

import (
	"fmt"
	"sumi/app/models"
	"sumi/app/store"
)

// ReviewsRepository ...
type ReviewsRepository struct {
	store *Store
}

func (r ReviewsRepository) Create(review *models.Review) (*models.Review, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO sumireviews (blogger_id, post_link) VALUES ($1, $2) RETURNING id",
		review.BloggerID, review.PostLink,
	).Scan(&review.ID); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return review, nil
}

func (r *ReviewsRepository) GetAll() ([]*models.Review, error) {

	var reviews []*models.Review
	rows, err := r.store.db.Query("SELECT * FROM sumireviews")
	if err != nil {
		return nil, err
	}
	reviews = ParseReviews(rows)
	return reviews, nil

}

func (r *ReviewsRepository) Delete() error {
	res, err := r.store.db.Exec("DELETE FROM sumireviews")
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	} else {
		return store.ErrNoRowsAffected
	}
}
