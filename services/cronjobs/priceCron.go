package cronjobs

import (
	"context"
	"invest/errors"
	"invest/services"

	"github.com/robfig/cron/v3"
)

func StartCronJobs() {
	cron := cron.New()
	cron.AddFunc("@every 1w", func() {
		priceCron()
	})
	cron.Start()
}

	func priceCron() error {
		_, err := services.GetAllAccounts(context.Background())
		if err != nil {
			return errors.Wrap(404, "account not found", err)
		}
	
		return nil
	}
