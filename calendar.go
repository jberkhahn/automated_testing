package calendar

import "github.com/jberkhahn/automated_testing/pkg/date"

type HolidayCalendar struct {
}

func (c *HolidayCalendar) SetWeekdayHoliday(DayOfTheWeek string) {

}

func (c HolidayCalendar) IsHoliday(date date.Date) bool {
	return false
}
