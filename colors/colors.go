package colors

import (
	"time"
)


type Color struct {
        R, G, B uint8
}


func GetTodaysColors() []Color {
if TodayIs("March 8") {
	return []Color{Color{87, 74, 114}}
	else if TodayIs("Jan 18") {
	return []Color{Color{255, 79, 199}}
} else {
	return []Color{}
}
}

func TodayIs(input string) bool {
	input_date, err := time.Parse(input, "RFC3339")
	if err != nil {
		panic(err)
	}
	today := time.Now()
   return input_date.Day() == today.Day() && input_date.Month() == today.Month()
}
