package date

import "time"

type Date struct {
	MyTime time.Time
}

func NewDate(year, month, day int) Date {
	return Date{
		MyTime: time.Date(year, time.Month(month), day, 12, 0, 0, 0, time.UTC),
	}
}

func (d Date) Year() int {
	return d.MyTime.Year()
}
func (d Date) Month() int {
	return int(d.MyTime.Month())
}
func (d Date) Day() int {
	return d.MyTime.Day()
}
func (d Date) Weekday() string {
	return d.MyTime.Weekday().String()
}
