package dto

type CreateUserDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUserDTO struct {
	Name *string `json:"name,omitempty"`
}
