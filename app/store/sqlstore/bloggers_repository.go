package sqlstore

import (
	"database/sql"
	"fmt"
	"sumi/app/models"
	"sumi/app/store"
)

// BloggersRepository ...
type BloggersRepository struct {
	store *Store
}

func (r *BloggersRepository) Create(b *models.Blogger) (*models.Blogger, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO sumibloggers (name, login, type, description, subs_count, avatar, social_network, cost, coverage) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		b.Name,
		b.Login,
		b.Type,
		b.Description,
		b.SubsCount,
		b.AvatarUrl,
		b.SocialNetwork,
		b.Cost,
		b.Coverage,
	).Scan(&b.ID); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return b, nil
}

func (r *BloggersRepository) GetAll() ([]*models.Blogger, error){

	var bloggers []*models.Blogger
		rows, err := r.store.db.Query("SELECT * FROM sumibloggers")
		if err != nil {
			return nil, err
		}
	return ParseBloggers(rows, bloggers)

}

func (r *BloggersRepository) GetByLogin(login string) (*models.Blogger, error){

	b := &models.Blogger{}
	if err :=  r.store.db.QueryRow("SELECT * FROM sumibloggers WHERE login = $1", login).Scan(
		&b.ID,
		&b.Name,
		&b.Login,
		&b.Description,
		&b.Cost,
		&b.SocialNetwork,
	); err != nil {
		return nil, err
	}
	return b, nil
}

func ParseBloggers(rows *sql.Rows, bloggers []*models.Blogger) ([]*models.Blogger, error) {
	for rows.Next() {
		b := models.Blogger{}

		err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.Login,
			&b.Type,
			&b.Description,
			&b.SubsCount,
			&b.AvatarUrl,
			&b.SocialNetwork,
			&b.Cost,
			&b.Coverage,
		)

		if err != nil {
			continue
		}

		bloggers = append(bloggers, &b)
	}

	return bloggers, nil
}

func (r *BloggersRepository) Delete() error {
	res, err := r.store.db.Exec("DELETE FROM sumibloggers")
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