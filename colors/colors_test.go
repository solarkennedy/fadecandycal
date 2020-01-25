package colors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, TodayIs("March 8", march8), true)
}

func TestTodayIsRange(t *testing.T) {
	const shortForm = "2006-Jan-02"
	march8, _ := time.Parse(shortForm, "2001-Mar-08")
	assert.Equal(t, TodayIsRange("March 6", 3, march8), true)
	assert.Equal(t, TodayIsRange("March 7", 3, march8), true)
	assert.Equal(t, TodayIsRange("March 8", 3, march8), true)
	assert.Equal(t, TodayIsRange("March 9", 3, march8), false)
	assert.Equal(t, TodayIsRange("March 10", 3, march8), false)
	assert.Equal(t, TodayIsRange("March 11", 3, march8), false)
}

func TestTodayIsRange2(t *testing.T) {
	const shortForm = "2006-Jan-02"
	dec24, _ := time.Parse(shortForm, "2020-Dec-24")
	assert.Equal(t, TodayIsRange("January 1", 1, dec24), false)
	assert.Equal(t, TodayIsRange("January 1", 31, dec24), false)
}
