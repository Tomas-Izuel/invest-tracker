package cronjobs

import (
	"github.com/robfig/cron/v3"
)

func StartCronJobs() {
	cron := cron.New()
	cron.AddFunc("@every 1w", func() {
		priceCron()
	})
	cron.Start()
}