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
	Logger.Info(">> Starting")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer close(c)

	initScheduler()
	initWebServer(c)

	Logger.Info(">> Started")
	select {
	case <-c:
		Logger.Info(">> Exiting")
		return
	}
}

func initWebServer(c chan os.Signal) {
	r := http.NewServeMux()
	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.HandleFunc("/api/v1/example", func(w http.ResponseWriter, r *http.Request) {

	})

	// ctrl + c == stop!!
	go func() {
		defer func() {
			// shut down
			c <- syscall.SIGTERM
		}()
		if err := http.ListenAndServe(":8080", r); err != nil {
			Logger.Error("ListenAndServe", zap.Error(err))
		}
	}()
}

func initScheduler() {
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
}

type CronJob func()

func (j CronJob) Run() {
	j()
}
