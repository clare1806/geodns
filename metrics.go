package main

import (
	metrics "github.com/abh/go-metrics"
	"log"
	"os"
	"runtime"
	"time"
)

var qCounter = metrics.NewMeter()

func metricsPoster() {

	lastQueryCount := qCounter.Count()

	queries := metrics.NewMeter()
	metrics.Register("queries", queries)

	queriesHistogram := metrics.NewHistogram(metrics.NewUniformSample(600))
	metrics.Register("queriesHistogram", queriesHistogram)

	goroutines := metrics.NewGauge()
	metrics.Register("goroutines", goroutines)

	go metrics.Log(metrics.DefaultRegistry, 30, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	// metrics.

	for {
		time.Sleep(1 * time.Second)

		// log.Println("updating metrics")

		current := qCounter.Count()
		newQueries := current - lastQueryCount
		lastQueryCount = current

		queries.Mark(newQueries)
		queriesHistogram.Update(newQueries)
		goroutines.Update(int64(runtime.NumGoroutine()))

	}
}