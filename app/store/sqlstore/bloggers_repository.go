package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"sumi/app/models"
	"sumi/app/store"
)

// BloggersRepository ...
type BloggersRepository struct {
	store *Store
}

func (r *BloggersRepository) Create(b *models.Blogger) (*models.Blogger, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO sumibloggers (name, login, type, description, subs_count, avatar, social_network, cost, coverage, is_selected, is_favorite) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		b.Name,
		b.Login,
		b.Type,
		b.Description,
		b.SubsCount,
		b.AvatarUrl,
		b.SocialNetwork,
		b.Cost,
		b.Coverage,
		b.IsSelected,
		false,
	).Scan(&b.ID); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return b, nil
}

func (r *BloggersRepository) GetAll(isSelected bool, isFavorite bool) ([]*models.Blogger, error){

	sqlQuery := "SELECT * FROM sumibloggers WHERE is_selected = $1 AND is_favorite = $2"

	var bloggers []*models.Blogger
		rows, err := r.store.db.Query(sqlQuery, isSelected, isFavorite)
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
		&b.IsSelected,
		&b.IsFavorite,
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
			&b.IsSelected,
			&b.IsFavorite,
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

func (r *BloggersRepository) Select(ids []string) error {
	idsQuery := fmt.Sprintf(" WHERE id in (%s)", strings.Join(ids, ", "))
	res, err := r.store.db.Exec("UPDATE sumibloggers SET is_selected = not is_selected" + idsQuery)
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