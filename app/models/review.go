package models

import validation "github.com/go-ozzo/ozzo-validation"

type Review struct {
	ID          int    `json:"id"`
	BloggerID   int    `json:"blogger_id"`
	PressTourId int    `json:"press_tour_id"`
	PostLink    string `json:"post_link"`
	Mark        int    `json:"mark"`
}

// Validate ...
func (r *Review) Validate() error {
	return validation.ValidateStruct(r, validation.Field(
		&r.BloggerID, validation.Required,
	), validation.Field(&r.PressTourId, validation.Required), validation.Field(&r.Mark, validation.Max(5)))
}
