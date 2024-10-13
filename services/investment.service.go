package services

import (
	"context"
	"invest/errors"
	"invest/lib"
	"invest/models"
	"invest/models/dto"
	"invest/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateInvestment(ctx context.Context, investmentDTO *dto.CreateInvestmentDTO) error {
	if err := validate.Struct(investmentDTO); err != nil {
		return errors.NewValidationError(lib.MapValidationErrors(err))
	}

	accountExist, err := repository.FindAccountByID(ctx, investmentDTO.AccountID)
	if err != nil || accountExist == nil {
		return errors.Wrap(404, "account not found", err)
	}

	investment := &models.Investment{
		AccountID: accountExist.ID,
		Name:      investmentDTO.Name,
		Code:      investmentDTO.Code,
		Stock:     investmentDTO.Stock,
	}

	if err := repository.InsertInvestment(ctx, investment); err != nil {
		return errors.Wrap(500, "failed to create investment", err)
	}

	return nil
}

func UpdateInvestment(ctx context.Context, id string, investmentDTO *dto.UpdateInvestmentDTO) error {
	if err := validate.Struct(investmentDTO); err != nil {
		return errors.NewValidationError(lib.MapValidationErrors(err))
	}

	investmentExist, err := repository.FindInvestmentByID(ctx, id)
	if err != nil || investmentExist == nil {
		return errors.Wrap(404, "investment not found", err)
	}

	updateData := map[string]interface{}{
		"name":  investmentDTO.Name,
		"code":  investmentDTO.Code,
		"stock": investmentDTO.Stock,
	}

	if err := repository.UpdateInvestment(ctx, id, updateData); err != nil {
		return errors.Wrap(500, "failed to update investment", err)
	}

	return nil
}

func DeleteInvestment(ctx context.Context, id string) error {
	investmentExist, err := repository.FindInvestmentByID(ctx, id)
	if err != nil || investmentExist == nil {
		return errors.Wrap(404, "investment not found", err)
	}

	if err := repository.DeleteInvestment(ctx, id); err != nil {
		return errors.Wrap(500, "failed to delete investment", err)
	}

	return nil
}

func GetInvestmentByID(ctx context.Context, id string) (*models.Investment, error) {
	investment, err := repository.FindInvestmentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return investment, nil
}

func InsertInvestmentPrice(ctx context.Context, investmentId string, total float64) error {
	if investmentExist, err := repository.FindInvestmentByID(ctx, investmentId); err != nil || investmentExist == nil {
		return errors.Wrap(404, "investment not found", err)
	}

	investmentID, err := primitive.ObjectIDFromHex(investmentId)
	if err != nil {
		return errors.Wrap(400, "invalid investment ID format", err)
	}

	investmentPrice := &models.InvestmentPrice{
		InvestmentID: investmentID,
		TotalPrice:   total,
		Date:         time.Now(),
	}

	if err := repository.InsertInvestmentPrice(ctx, investmentPrice); err != nil {
		return errors.Wrap(500, "failed to insert investment price", err)
	}

	return nil
}
