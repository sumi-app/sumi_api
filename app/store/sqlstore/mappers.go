package sqlstore

import (
	"database/sql"
	"sumi/app/models"
)

func ParseReviews(rows *sql.Rows) []*models.Review {
	var reviews []*models.Review
	for rows.Next() {
		r := models.Review{}

		err := rows.Scan(
			&r.ID,
			&r.BloggerID,
			&r.PostLink,
			&r.Likes,
		)

		if err != nil {
			continue
		}

		reviews = append(reviews, &r)
	}
	return reviews
}