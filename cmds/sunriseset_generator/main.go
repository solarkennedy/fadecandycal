package main

import (
	"fmt"
	"github.com/kelvins/sunrisesunset"
	"strconv"
	"time"
)

func getUTCOffset(now time.Time) float64 {
	offset, err := strconv.Atoi(now.Format("-0700"))
	if err != nil {
		panic(err)
	}
	return float64(offset / 100)
}

func getSunriseSunset(now time.Time) (time.Time, time.Time) {
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	p := sunrisesunset.Parameters{
		Latitude:  37.774929,
		Longitude: -122.419418,
		UtcOffset: getUTCOffset(now),
		Date:      today,
	}
	sunrise, sunset, err := p.GetSunriseSunset()
	if err != nil {
		panic(err)
	}
	sunrise_today := time.Date(now.Year(), now.Month(), now.Day(), sunrise.Hour(), sunrise.Minute(), sunrise.Second(), 0, now.Location())
	sunset_today := time.Date(now.Year(), now.Month(), now.Day(), sunset.Hour(), sunset.Minute(), sunset.Second(), 0, now.Location())
	fmt.Println(" Sunrise:", sunrise_today.Format("3:04PM"), " / Sunset:", sunset_today.Format("3:04PM"))
	return sunrise_today, sunset_today
}

func main() {

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	start, err := time.ParseInLocation("2006-1-2", "2020-01-01", loc)
	if err != nil {
		panic(err)
	}

	for d := start; d.Year() <= 2030; d = d.AddDate(0, 0, 1) {
		fmt.Print(d)
		fmt.Print(": ")
		getSunriseSunset(d)
	}
}
