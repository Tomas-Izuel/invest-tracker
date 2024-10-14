package dto

type CreateAccountTypeDTO struct {
	Name string `json:"name" validate:"required"`
}
