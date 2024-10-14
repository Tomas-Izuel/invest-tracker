package services

import (
	"context"
	"invest/errors"
	"invest/lib"
	"invest/models"
	"invest/models/dto"
	"invest/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAccountType(ctx context.Context, accountTypeDTO *dto.CreateAccountTypeDTO) (*mongo.InsertOneResult, error) {
	if err := validate.Struct(accountTypeDTO); err != nil {
		return nil, errors.NewValidationError(lib.MapValidationErrors(err))
	}

	accountType := &models.AccountType{
		Name: accountTypeDTO.Name,
	}

	created, err := repository.InsertAccountType(ctx, accountType)
	if err != nil {
		return nil, errors.Wrap(500, "failed to create account type", err)
	}

	return created, nil
}

func GetAllAccountTypes(ctx context.Context) ([]models.AccountType, error) {
	accountTypes, err := repository.FindAccountTypes(ctx)
	if err != nil {
		return nil, err
	}

	return accountTypes, nil
}

func UpdateAccountType(ctx context.Context, id string, accountTypeDTO *dto.CreateAccountTypeDTO) (*mongo.UpdateResult, error) {
	if err := validate.Struct(accountTypeDTO); err != nil {
		return nil, errors.NewValidationError(lib.MapValidationErrors(err))
	}

	accountType := &models.AccountType{
		Name: accountTypeDTO.Name,
	}

	updated, err := repository.UpdateAccountType(ctx, id, accountType)
	if err != nil {
		return nil, errors.Wrap(500, "failed to update account type", err)
	}

	return updated, nil
}
