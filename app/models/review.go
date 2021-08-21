package models

import validation "github.com/go-ozzo/ozzo-validation"

type Review struct {
	ID        int    `json:"id"`
	BloggerID int    `json:"blogger_id"`
	PostLink  string `json:"post_link"`
}

// Validate ...
func (r *Review) Validate() error {
	return validation.ValidateStruct(r, validation.Field(
		&r.BloggerID, validation.Required,
	))
}
