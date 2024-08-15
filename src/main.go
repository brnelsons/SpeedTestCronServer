package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/calin014/speedfast"
	"go.uber.org/zap"
)

var Logger = zap.Must(zap.NewProduction())

func main() {
	scheduler := cron.New()
	defer scheduler.Stop()
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	go func() {
		if err := http.ListenAndServe(":8080", mux); err != nil {
			Logger.Error("ListenAndServe", zap.Error(err))
		}
	}()

	// ctrl + c == stop!!
	c := make(chan os.Signal)
	defer close(c)
	go func() {
		defer func() {
			// shut down
			c <- syscall.SIGTERM
		}()
		if err := http.ListenAndServe(":8080", mux); err != nil {
			Logger.Error("ListenAndServe", zap.Error(err))
		}
	}()
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	select {
	case <-c:
		return
	}
}

type CronJob func()

func (j CronJob) Run() {
	j()
}
