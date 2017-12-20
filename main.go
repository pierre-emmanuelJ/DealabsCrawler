package main

import (
	"time"

	dealcrawler "github.com/pierre-emmanuelJ/DealabsCrawler/dealabsCrawler"
)

func main() {
	dealcrawler.AllComment = nil
	for {
		dealcrawler.Crawler()
		//TODO put cron
		time.Sleep(10 * time.Second)
	}
}
