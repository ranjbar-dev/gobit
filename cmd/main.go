package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ranjbar-dev/gobit/srv/api"
	"github.com/ranjbar-dev/gobit/srv/cron"
)

func main() {

	defer func() {

		if err := recover(); err != nil {

			// golog.Logger.Error("Error in engine", "recover function called", err)
		}
	}()

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	waitChannel := make(chan struct{}, 1)
	go func() {

		// we can exit from app now
		defer func() {

			waitChannel <- struct{}{}
		}()

		// wait for signal to exit
		<-sigs
		cancel()
	}()

	// start api
	a := api.NewApi(ctx, cancel)
	a.Start()

	// cron jobs
	cron.NewCronService(ctx).Start()

	// wait to exit from app
	<-waitChannel
}
