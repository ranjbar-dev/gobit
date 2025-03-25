package cron

import (
	"context"
	"time"
)

type CronJob struct {
	ctx      context.Context
	ticker   *time.Ticker
	callback func()
}

func (j *CronJob) Start() {

	for {

		select {

		case <-j.ctx.Done():

			j.ticker.Stop()
			return

		case <-j.ticker.C:

			j.callback()
		}

	}
}
