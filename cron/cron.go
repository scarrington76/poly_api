package cron

import (
	"api/polygon/stocks"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"golang.org/x/text/number"
)

func Start(h int64, m int64) {

	s1 := gocron.NewScheduler(time.UTC)

	n := number.Decimal(
		h, number.Pad('0'), number.FormatWidth(2),
	)
	print(n)
	currentday := time.Now().Format("2006-01-02")

	_, err := s1.Every(1).Day().At("00:01").Do(func() { stocks.AllStocksDaily(currentday) })
	if err != nil {
		log.Fatalf("error creating job: %v", err)
	}

}
