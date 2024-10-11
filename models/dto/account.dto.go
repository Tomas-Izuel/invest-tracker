package dto

type CreateAccountDTO struct {
	Name        string `json:"name" validate:"required"`
	Period      string `json:"period" validate:"required"`
	EndpointURL string `json:"endpoint_url" validate:"required"`
	HasLogin    bool   `json:"has_login"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	UserID      string `json:"user_id" validate:"required"` // Referencia a User
}

type UpdateAccountDTO struct {
	Name        *string `json:"name,omitempty"`
	Period      *string `json:"period,omitempty"`
	EndpointURL *string `json:"endpoint_url,omitempty"`
	HasLogin    *bool   `json:"has_login,omitempty"`
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`
}

type DeleteAccountDTO struct {
	AccountID string `json:"account_id" validate:"required"`
}
