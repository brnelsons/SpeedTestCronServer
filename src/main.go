package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/calin014/speedfast"
)

func main() {
	scheduler := cron.New()
	// minute (0-59), hour (0-23), day of the month (1-31), month (1-12), day of the week (0-6)
	_, err := scheduler.AddJob("*/1 * * * *", CronJob(func() {
		s := time.Now()
		fmt.Println(speedfast.MeasureWithFast())
		fmt.Println("took: ", time.Since(s))
	}))
	if err != nil {
		panic(err)
	}

	scheduler.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	select {
	case <-c:
		scheduler.Stop()
	}
}

type CronJob func()

func (j CronJob) Run() {
	j()
}
