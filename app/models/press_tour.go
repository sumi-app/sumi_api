package models

type PressTour struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	CommonMark  int    `json:"common_mark"`
}
