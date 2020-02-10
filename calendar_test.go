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

func TestWeekdayNotHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 6, 1)

	if c.IsHoliday(d) == true {
		t.Errorf("Monday should not be a holiday")
	}
}

func TestMultipleWeekdayHolidays(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 6, 6)
	d2 := date.NewDate(2020, 6, 7)

	c.SetWeekdayHoliday("Saturday")
	c.SetWeekdayHoliday("Sunday")

	if c.IsHoliday(d) != true {
		t.Errorf("Saturday should be a holiday")
	}
	if c.IsHoliday(d2) != true {
		t.Errorf("Sunday should be a holiday")
	}
}

func TestSpecificHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 1, 1)

	c.SetDateHoliday(2020, 1, 1)

	if c.IsHoliday(d) != true {
		t.Errorf("New Years Day should be a holiday")
	}
}
