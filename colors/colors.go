package colors

import (
	"fmt"
	"github.com/gookit/color"
	"strconv"
	"strings"
	"time"
)

type Color struct {
	R, G, B uint8
}

func PrintColors(colors []Color, occasion string, day time.Time) {
	fmt.Printf("For %s (%s) our colors will be [", occasion, day)
	for _, c := range colors {
		cprint := color.RGB(c.R, c.G, c.B) // fg color
		cprint.Print("█")
	}
	fmt.Println("]")
}

func GetDaysColors(day time.Time) []Color {
	colors := []Color{}
	occasion := ""

	if TodayIs("January 1", day) {
		occasion = "New Years Day"
		colors = []Color{Color{212, 175, 55}, Color{0, 0, 0}}
	} else if TodayIs("January 18", day) {
		occasion = "Woman's March"
		colors = []Color{Color{255, 79, 199}, Color{0, 0, 0}}
	} else if TodayIs("January 20", day) {
		occasion = "MLK Day"
		colors = []Color{Color{239, 52, 35}, Color{255, 209, 2}, Color{46, 151, 67}, Color{}}
	} else if TodayIs("January 25", day) {
		occasion = "Chinese New Year"
		colors = []Color{Color{120, 0, 0}, Color{168, 0, 0}, Color{190, 24, 0}, Color{255, 212, 37}, Color{239, 196, 22}}
	} else if TodayIs("January 29", day) {
		occasion = "NEN Awards"
		colors = []Color{Color{151, 193, 209}, Color{123, 170, 124}, Color{250, 187, 24}}
	} else if TodayIsRange("February 1", 2, day) {
		occasion = "Red/Gold - 49ers at Superbowl"
		colors = []Color{Color{187, 52, 49}, Color{186, 150, 106}}
	} else if TodayIs("February 4", day) {
		occasion = "Emperor Norton's 200th Birthday"
		colors = []Color{Color{212, 175, 55}, Color{0, 0, 0}}
	} else if TodayIsRange("February 19", 5, day) {
		occasion = "Black History Month"
		colors = []Color{Color{239, 52, 35}, Color{255, 209, 2}, Color{46, 151, 67}, Color{}}
	} else if TodayIs("March 8", day) {
		occasion = "International Women's Day"
		colors = []Color{Color{87, 74, 114}, Color{0, 0, 0}}
	} else {
		occasion = "(No Occasion)"
		colors = []Color{}
	}

	PrintColors(colors, occasion, day)
	return colors
}

func TodayIs(input string, today time.Time) bool {
	s := strings.Split(input, " ")
	month := s[0]
	day, _ := strconv.Atoi(s[1])
	return day == today.Day() && month == today.Month().String()
}

func TodayIsRange(input string, n int, today time.Time) bool {
	//BUG: Normalizes the date comparison within the same year,
	//so it can't really span year boundaries
	//	fmt.Printf("Is %s within %d days after %s?\n", today, n, input)
	input_date := parse_input_date(input, today)
	last_date := input_date.AddDate(0, 0, n)
	//	fmt.Printf("(Between %s and %s\n", input_date, last_date)
	result := (today.After(input_date) && today.Before(last_date)) || input_date == today
	//	fmt.Println(result)
	return result
}

func parse_input_date(input string, normalized_day time.Time) time.Time {
	s := strings.Split(input, " ")
	month := s[0]
	day, _ := strconv.Atoi(s[1])
	parsed := time.Date(normalized_day.Year(), MonthToMonth(month), day, 0, 0, 0, 0, normalized_day.Location())
	return parsed
}

func MonthToMonth(input string) time.Month {
	fake_date := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Now().Location())
	month := fake_date.Month()
	for i := 1; i < 12; i++ {
		if fake_date.Month().String() == input {
			return fake_date.Month()
		}
		fake_date = fake_date.AddDate(0, 1, 0)

	}
	panic(fmt.Sprintf("What is this month? %s", input))
	return month
}
