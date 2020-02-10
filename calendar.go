package calendar

import "github.com/jberkhahn/automated_testing/pkg/date"

type HolidayCalendar struct {
	weekdayHolidays  []string
	specificHolidays []date.Date
}

func (c *HolidayCalendar) SetWeekdayHoliday(DayOfTheWeek string) {
	c.weekdayHolidays = append(c.weekdayHolidays, DayOfTheWeek)
}

func (c *HolidayCalendar) SetDateHoliday(year, month, day int) {
	c.specificHolidays = append(c.specificHolidays, date.NewDate(year, month, day))
}

func (c *HolidayCalendar) IsHoliday(inputDate date.Date) bool {
	return c.isWeekdayHoliday(inputDate) || c.isSpecificHoliday(inputDate)
}

func (c *HolidayCalendar) isWeekdayHoliday(inputDate date.Date) bool {
	for _, v := range c.weekdayHolidays {
		if v == inputDate.Weekday() {
			return true
		}
	}
	return false
}

func (c *HolidayCalendar) isSpecificHoliday(inputDate date.Date) bool {
	for _, v := range c.specificHolidays {
		if v == inputDate {
			return true
		}
	}
	return false
}
