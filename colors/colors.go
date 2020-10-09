package colors

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
)

type Color struct {
	R, G, B uint8
}

func PrintColors(colors []Color, occasion string, day time.Time) {
	fmt.Printf("For %s (%s) our colors will be [", occasion, day)
	for _, c := range colors {
		PrintColorBlock(c)
	}
	fmt.Println("]")
}

func PrintColorBlock(c Color) {
	cprint := color.RGB(c.R, c.G, c.B) // fg color
	cprint.Print("█")
}

func GetDaysColors(day time.Time) []Color {
	var colors []Color
	occasion := ""

	if TodayIs("January 1", day) {
		occasion = "New Years Day"
		colors = []Color{{212, 175, 55}, {0, 0, 0}}
	} else if TodayIs("January 18", day) {
		occasion = "Woman's March"
		colors = []Color{{255, 79, 199}, {0, 0, 0}}
	} else if TodayIs("January 20", day) {
		occasion = "MLK Day"
		colors = []Color{{239, 52, 35}, {255, 209, 2}, {46, 151, 67}, {}}
	} else if TodayIs("January 25", day) {
		occasion = "Chinese New Year"
		colors = []Color{
			{0, 0, 0},
			{255, 0, 0},
			{190, 24, 0},
			{255, 212, 37},
			{239, 196, 22},
		}
	} else if TodayIs("January 29", day) {
		occasion = "NEN Awards"
		colors = []Color{{151, 193, 209}, {123, 170, 124}, {250, 187, 24}}
	} else if TodayIsRange("January 1", 31, day) {
		occasion = "January"
		colors = []Color{
			{13, 10, 94},
			{125, 123, 144},
			{60, 141, 135},
			{109, 65, 139},
			{124, 67, 105},
			{255, 255, 255},
		}
	} else if TodayIsRange("February 1", 2, day) {
		occasion = "Red/Gold - 49ers at Superbowl"
		colors = []Color{{187, 52, 49}, {186, 150, 106}}
	} else if TodayIs("February 4", day) {
		occasion = "Emperor Norton's 200th Birthday"
		colors = []Color{{212, 175, 55}, {0, 0, 0}}
	} else if TodayIsRange("February 19", 5, day) {
		occasion = "Black History Month"
		colors = []Color{{239, 52, 35}, {255, 209, 2}, {46, 151, 67}, {}}
	} else if TodayIsRange("February 1", 28, day) {
		occasion = "February"
		colors = []Color{
			{236, 156, 192},
			{240, 177, 212},
			{246, 207, 245},
			{227, 219, 255},
			{210, 213, 253},
		}
	} else if TodayIs("March 8", day) {
		occasion = "International Women's Day"
		colors = []Color{{87, 74, 114}, {0, 0, 0}}
	} else if TodayIs("March 17", day) {
		occasion = "Saint Patrick's Day"
		colors = []Color{{0, 153, 89}, {0, 0, 0}}
	} else if TodayIsRange("March 19", 2, day) {
		occasion = "Nowruz/Persian New Year"
		colors = []Color{
			{255, 0, 0},
			{0, 255, 0},
			{255, 255, 255},
			{0, 0, 0},
		}
	} else if TodayIs("March 21", day) {
		occasion = "American Red Cross Day"
		colors = []Color{{255, 0, 0}, {0, 0, 0}}
	} else if TodayIs("March 24", day) {
		occasion = "World TB Day"
		colors = []Color{{153, 0, 0}, {0, 0, 0}}
	} else if TodayIs("March 25", day) {
		occasion = "National Cerebral Palsy Awareness Month"
		colors = []Color{{31, 158, 31}, {0, 0, 0}}
	} else if TodayIs("March 26", day) {
		occasion = "Colon Cancer Awareness Month"
		colors = []Color{{0, 128, 255}, {0, 0, 0}, {255, 255, 255}}
	} else if TodayIsRange("March 1", 31, day) {
		occasion = "March"
		colors = []Color{
			{111, 139, 199},
			{84, 180, 149},
			{74, 146, 90},
			{69, 141, 53},
			{63, 75, 133},
			{0, 0, 0},
		}
	} else if TodayIsRange("April 1", 30, day) {
		occasion = "April"
		colors = []Color{
			{184, 149, 214},
			{197, 181, 229},
			{206, 216, 255},
			{232, 255, 210},
			{248, 249, 171},
		}
	} else if TodayIsRange("May 1", 31, day) {
		occasion = "May"
		colors = []Color{
			{27, 236, 123},
			{02, 246, 139},
			{49, 255, 150},
			{55, 204, 222},
			{34, 203, 249},
		}
	} else if TodayIsRange("June 1", 30, day) {
		occasion = "June"
		colors = []Color{
			{255, 239, 63},
			{112, 224, 255},
			{227, 160, 242},
			{255, 154, 219},
			{204, 255, 0},
		}
	} else if TodayIsRange("July 1", 31, day) {
		occasion = "July"
		colors = []Color{
			{44, 201, 80},
			{38, 78, 90},
			{2, 154, 201},
			{5, 186, 125},
			{51, 124, 84},
		}
	} else if TodayIs("August 8", day) {
		occasion = "Pantone 448c"
		colors = []Color{{74, 65, 42}, {}}
	} else if TodayIsRange("August 1", 31, day) {
		occasion = "August"
		colors = []Color{

			{2, 135, 188},
			{9, 103, 167},
			{32, 93, 146},
			{54, 204, 109},
			{39, 131, 66},
		}
		//	} else if TodayIsRange("September 1", 30, day) {
		//		occasion = "September"
		//		colors = []Color{
		//			{78, 176, 129},
		//			{4, 193, 186},
		//			{8, 129, 85},
		//			{93, 170, 139},
		//			{03, 152, 211},
		//		}
	} else if TodayIsRange("October 1", 31, day) {
		occasion = "October"
		colors = []Color{
			{47, 141, 0},
			{20, 80, 0},
			{01, 89, 5},
			{8, 23, 3},
			{7, 0, 0},
		}
	} else if TodayIsRange("November 1", 30, day) {
		occasion = "November"
		colors = []Color{
			{55, 195, 31},
			{20, 117, 0},
			{72, 67, 12},
			{18, 41, 0},
			{30, 18, 16},
		}
	} else if TodayIsRange("December 1", 31, day) {
		occasion = "December"
		colors = []Color{
			{87, 22, 20},
			{29, 0, 0},
			{8, 99, 24},
			{9, 46, 5},
			{5, 59, 46},
		}
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
	for i := 1; i <= 12; i++ {
		if fake_date.Month().String() == input {
			return fake_date.Month()
		}
		fake_date = fake_date.AddDate(0, 1, 0)

	}
	return time.Month(1)
}
