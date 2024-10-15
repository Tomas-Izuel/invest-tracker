package dto

type CreateInvestmentDTO struct {
	Name  string `json:"name" validate:"required"`
	Code  string `json:"code" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}

type UpdateInvestmentDTO struct {
	Name  *string `json:"name,omitempty"`
	Code  *string `json:"code,omitempty"`
	Stock *int    `json:"stock,omitempty"`
}

type DeleteInvestmentDTO struct {
	InvestmentID string `json:"investment_id" validate:"required"`
}
