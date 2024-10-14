package cronjobs

import (
	"context"
	"invest/builders"
	"invest/errors"
	"invest/services"
)

func priceCron() error {
		account, err := services.GetAllAccounts(context.TODO())
		if err != nil {
			return errors.Wrap(404, "account not found", err)
		}
	
		for _, ac := range account {
			builders.PriceBuilderMethod("")
			_ = ac
		}

		return nil
	}
