package sqlstore

import (
	"fmt"
	"sumi/app/models"
	"sumi/app/store"
)

// PressTourRepository ...
type PressTourRepository struct {
	store *Store
}

func (r PressTourRepository) Create(p *models.PressTour) (*models.PressTour, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO sumipresstours (title, description, image_url) VALUES ($1, $2, $3) RETURNING id",
		p.Title, p.Description, p.ImageUrl,
	).Scan(&p.ID); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return p, nil
}


func (r *PressTourRepository) GetAll() ([]*models.PressTour, error){

	var pressTours []*models.PressTour
	rows, err := r.store.db.Query("SELECT * FROM sumipresstours")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := models.PressTour{}

		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.ImageUrl,
		)

		if err != nil {
			continue
		}

		pressTours = append(pressTours, &p)
	}

	return pressTours, nil

}


func (r *PressTourRepository) Delete() error {
	res, err := r.store.db.Exec("DELETE FROM sumipresstours")
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
