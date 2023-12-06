package main

import (
	"price-comparison/amazon"

	"time"

	"github.com/go-co-op/gocron"
)

func cronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(60).Seconds().Do(func() {
		amazon.AmazonProduct()
	})

	s.StartBlocking()
}

func main() {

	cronJobs()
}
