package cron

import (
	"context"
)

type Cron struct {
	ctx  context.Context
	jobs []CronJob
}

func (c *Cron) registerCronJobs() *Cron {

	// delete old quotes
	// c.addJob(CronJob{
	// 	ctx:    c.ctx,
	// 	ticker: time.NewTicker(time.Hour * 6),
	// 	callback: func() {

	// 		from := timetool.Now().Add(-time.Hour * 24 * 7).UnixMilli()
	// 		err := influxrepository.NewInfluxRepository(context.Background()).DeleteQuotesBefore(from)
	// 		if err != nil {

	// 			golog.Logger.Error("error in deleting old quotes in Cron service", err.Error())
	// 		} else {

	// 			golog.Logger.Trace("delete quites in Cron service", "old quotes deleted successfully")
	// 		}
	// 	},
	// })

	// sync charts
	// c.addJob(CronJob{
	// 	ctx:    c.ctx,
	// 	ticker: time.NewTicker(time.Hour * 6),
	// 	callback: func() {

	// 		err := chart.NewChartService().SyncCharts()
	// 		if err != nil {

	// 			golog.Logger.Error("error in syncing charts in Cron service", err.Error())
	// 		} else {

	// 			golog.Logger.Trace("chart sync in Cron service", "charts sycned successfully")
	// 		}
	// 	},
	// })

	// delete guest accounts
	// c.addJob(CronJob{
	// 	ctx:    c.ctx,
	// 	ticker: time.NewTicker(time.Hour * 24),
	// 	callback: func() {

	// 		// get all guest accounts
	// 		accounts, err := c.e.AllGuestAccounts()
	// 		if err != nil {

	// 			golog.Logger.Error("error in syncing charts in Cron service", err.Error())
	// 			return
	// 		}

	// 		// delete accounts created before 7 days
	// 		expirationTime := timetool.Now().Add(-time.Hour * 24 * 7).UnixMilli()

	// 		for _, account := range accounts {

	// 			if account.CreatedAt < expirationTime {

	// 				_, err := c.e.DeleteAccount(account.ID)
	// 				if err != nil {

	// 					golog.Logger.Error("error in deleting guest account in Cron service", err.Error())
	// 				} else {

	// 					golog.Logger.Trace(fmt.Sprintf("delete guest account %d in Cron service", account.ID), "guest account deleted successfully", map[string]any{"account_id": account.ID})
	// 				}
	// 			}
	// 		}

	// 	},
	// })

	return c
}

func (c *Cron) addJob(job CronJob) {

	c.jobs = append(c.jobs, job)
}

func (c *Cron) Start() {

	for _, job := range c.jobs {

		go job.Start()
	}
}

func NewCronService(ctx context.Context) *Cron {

	return (&Cron{
		ctx,
		[]CronJob{},
	}).registerCronJobs()
}
