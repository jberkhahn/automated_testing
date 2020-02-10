package calendar

import (
	"testing"

	"github.com/jberkhahn/automated_testing/pkg/date"
)

func TestWeekdayHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 6, 6)

	c.SetWeekdayHoliday("Saturday")

	if c.IsHoliday(d) != true {
		t.Errorf("Saturday should be a holiday")
	}
}
