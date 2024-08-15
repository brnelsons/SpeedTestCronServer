package main

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/calin014/speedfast"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var Logger = zap.Must(zap.NewProduction())
var historyEntries = make([]HistoryEntry, 0)

func main() {
	Logger.Info(">> Starting")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer close(c)

	sched := initScheduler()
	defer sched.Stop()
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
	r.HandleFunc("/api/v1/history", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j := json.NewEncoder(w)
		err := j.Encode(historyEntries)
		if err != nil {
			Logger.Error("error encoding json", zap.Error(err))
		}
	})
	r.Handle("/scripts", http.FileServer(http.Dir("./static/scripts/")))
	r.Handle("/", http.FileServer(http.Dir("./static/")))

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

func initScheduler() *cron.Cron {
	scheduler := cron.New()
	// minute (0-59), hour (0-23), day of the month (1-31), month (1-12), day of the week (0-6)
	_, err := scheduler.AddJob("* * * * *", CronJob(func() {
		s := time.Now()
		measurement, err := speedfast.MeasureWithFast()
		if err != nil {
			Logger.Error("measurement", zap.Error(err))
		}
		historyEntries = append(historyEntries, HistoryEntry{
			Time:            s.UnixMilli(),
			Upload:          measurement.Upload,
			Download:        measurement.Download,
			MeasurementTime: time.Since(s).Milliseconds(),
		})
	}))
	if err != nil {
		panic(err)
	}
	scheduler.Start()
	return scheduler
}

type CronJob func()

func (j CronJob) Run() {
	j()
}

type HistoryEntry struct {
	Time            int64   `json:"time"`
	Upload          float64 `json:"upload"`
	Download        float64 `json:"download"`
	MeasurementTime int64   `json:"measurementTime"`
}
