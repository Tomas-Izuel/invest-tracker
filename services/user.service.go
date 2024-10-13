package services

import (
	"context"
	"invest/errors"
	"invest/lib"
	"invest/models"
	"invest/models/dto"
	"invest/repository"
)

// CreateUser creates a new user using the provided DTO
func CreateUser(ctx context.Context, userDTO *dto.CreateUserDTO) error {
	if err := validate.Struct(userDTO); err != nil {
		return errors.NewValidationError(lib.MapValidationErrors(err))
	}

	userExist, err := repository.FindUserByName(ctx, userDTO.Name)

	if err != nil || userExist != nil {
		return errors.Wrap(400, "user already exists", err)
	}

	user := &models.User{
		Name: userDTO.Name,
	}

	if err := repository.InsertUser(ctx, user); err != nil {
		return errors.Wrap(500, "failed to create user", err)
	}

	return nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := repository.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
