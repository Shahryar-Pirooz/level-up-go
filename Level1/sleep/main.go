package main

import (
	"flag"
	"log"
	"time"
)

const expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	time, err := time.Parse(expectedFormat, target)
	if err != nil {
		log.Fatal(err)
	}
	return time
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	now := time.Now()
	sleeps := target.Sub(now).Hours() / 24
	return float64(sleeps)
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
