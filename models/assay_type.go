package models

type AssayTypeInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
