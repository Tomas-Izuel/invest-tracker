package dto

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
