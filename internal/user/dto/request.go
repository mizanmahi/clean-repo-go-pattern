package dto

type CreateRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UpdateRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}
