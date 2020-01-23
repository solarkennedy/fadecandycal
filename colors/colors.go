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

func PrintColors(colors []Color, occasion string) {
	fmt.Printf("For %s our colors will be [", occasion)
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
		colors = []Color{Color{239, 52, 35}, Color{255, 209,  2}, Color{ 46, 151 , 67}, Color{}}
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
		colors = []Color{Color{239, 52, 35}, Color{255, 209,  2}, Color{ 46, 151 , 67}, Color{}}
	} else if TodayIs("March 8", day) {
		occasion = "International Women's Day"
		colors = []Color{Color{87, 74, 114}, Color{0, 0, 0}}
	} else {
		occasion = "Unknown"
		colors = []Color{}
	}

	PrintColors(colors, occasion)
	return colors
}

func TodayIs(input string, today time.Time) bool {
	s := strings.Split(input, " ")
	month := s[0]
	day, _ := strconv.Atoi(s[1])
	return day == today.Day() && month == today.Month().String()
}

func TodayIsRange(input string, n int, today time.Time) bool {
	for i := 0; i < n; i++ {
		if TodayIs(input, today.AddDate(0, 0, i)) {
			return true
		}
	}
	return false
}
