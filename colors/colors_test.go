package colors

import (
	"fmt"
	"testing"
	"time"
)

func TestAllDays(t *testing.T) {

	date := time.Now()
	for i := 0; i <= 365; i++ {
		testdate := date.AddDate(0, 0, i)
		fmt.Print("On ", testdate.Format("Mon 2006-01-2"), ": ")
		GetDaysColors(testdate)
	}
}

func TestTodayIs(t *testing.T) {
	const shortForm = "2006-Jan-02"
	march8, _ := time.Parse(shortForm, "2001-Mar-08")
	if TodayIs("March 8", march8) != true {
		t.Errorf("Wrong parsing for: march8")
	}
}
