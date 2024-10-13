package services

import (
	"context"
	"invest/errors"
	"invest/lib"
	"invest/models"
	"invest/models/dto"
	"invest/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllAcountsByUserID(ctx context.Context, userID string) ([]models.Account, error) {
	accounts, err := repository.GetAllAccountsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func CreateAccount(ctx context.Context, accountDTO *dto.CreateAccountDTO) error {
	if err := validate.Struct(accountDTO); err != nil {
		return errors.NewValidationError(lib.MapValidationErrors(err))
	}

	userID, err := primitive.ObjectIDFromHex(accountDTO.UserID)
	if err != nil {
		return errors.Wrap(400, "invalid user ID format", err)
	}

	userExist, err := repository.FindUserByID(ctx, accountDTO.UserID)
	if err != nil || userExist == nil {
		return errors.Wrap(404, "user not found", err)
	}

	account := &models.Account{
		UserID: userID,
		Name:   accountDTO.Name,
	}

	if err := repository.InsertAccount(ctx, account); err != nil {
		return errors.Wrap(500, "failed to create account", err)
	}

	return nil
}

func UpdateAccount(ctx context.Context, id string, accountDTO *dto.UpdateAccountDTO) error {
	if err := validate.Struct(accountDTO); err != nil {
		return errors.NewValidationError(lib.MapValidationErrors(err))
	}

	accountExist, err := repository.FindAccountByID(ctx, id)
	if err != nil || accountExist == nil {
		return errors.Wrap(404, "account not found", err)
	}

	updateData := map[string]interface{}{
		"name": accountDTO.Name,
	}

	if err := repository.UpdateAccount(ctx, id, updateData); err != nil {
		return errors.Wrap(500, "failed to update account", err)
	}

	return nil
}

func DeleteAccount(ctx context.Context, id string) error {
	accountExist, err := repository.FindAccountByID(ctx, id)
	if err != nil || accountExist == nil {
		return errors.Wrap(404, "account not found", err)
	}

	if err := repository.DeleteAccount(ctx, id); err != nil {
		return errors.Wrap(500, "failed to delete account", err)
	}

	return nil
}

func GetAccountByID(ctx context.Context, id string) (*models.Account, error) {
	account, err := repository.FindAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return account, nil
}
