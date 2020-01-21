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
	} else if TodayIs("January 20", day) {
		occasion = "MLK Day"
		colors = []Color{Color{}}
	} else if TodayIs("March 8", day) {
		occasion = "International Women's Day"
		colors = []Color{Color{87, 74, 114}, Color{0, 0, 0}}
	} else if TodayIs("January 18", day) {
		occasion = "Woman's March"
		colors = []Color{Color{255, 79, 199}, Color{0, 0, 0}}
	} else if TodayIs("January 29", day) {
		occasion = "NEN Awards"
		colors = []Color{Color{151, 193, 209}, Color{123, 170, 124}, Color{250, 187, 24}}
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
