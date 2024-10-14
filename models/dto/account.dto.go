package dto

import "invest/models"

type CreateAccountDTO struct {
	Period string `json:"period" validate:"required"`
	TypeID string `json:"type_id" validate:"required"`
}

type UpdateAccountDTO struct {
	Period *string `json:"period,omitempty"`
}

type DeleteAccountDTO struct {
	AccountID string `json:"account_id" validate:"required"`
}

type AccountResponseDTO struct {
	ID          string              `bson:"_id,omitempty"`
	Period      string              `json:"period"`
	UserID      string              `bson:"user_id"`
	Type        models.AccountType  `bson:"type"`
	Investments []models.Investment `bson:"investments,omitempty"`
}
