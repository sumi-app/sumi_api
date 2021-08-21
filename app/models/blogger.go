package models

import validation "github.com/go-ozzo/ozzo-validation"

type Blogger struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Login         string `json:"login"`
	Type          string `json:"type"`
	SocialNetwork int    `json:"social_network"`
	Description   string `json:"description"`
	SubsCount     int    `json:"count"`
	AvatarUrl     string `json:"avatar"`
	Cost          int    `json:"cost"`
	Coverage      int    `json:"coverage"`
	IsSelected    bool   `json:"is_selected"`
}

// Validate ...
func (b *Blogger) Validate() error {
	return validation.ValidateStruct(b, validation.Field(
		&b.Login, validation.Required,
	))
}
