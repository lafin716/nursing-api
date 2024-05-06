package date

import "time"

type formatter struct {
}

type Formatter interface {
	Format(t time.Time, layout string) string
	TruncateToDate(dateTime time.Time) time.Time
	TruncateToDateAddDay(dateTime time.Time, day int) time.Time
}

func NewFormatter() Formatter {
	return &formatter{}
}

func (f formatter) Format(t time.Time, layout string) string {
	replaced := ReplaceLayout(layout)
	return t.Format(replaced)
}

func (f formatter) TruncateToDate(dateTime time.Time) time.Time {
	return time.Date(
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		0,
		0,
		0,
		0,
		dateTime.Location(),
	)
}

func (f formatter) TruncateToDateAddDay(dateTime time.Time, day int) time.Time {
	return time.Date(
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day()+day,
		0,
		0,
		0,
		0,
		dateTime.Location(),
	)
}
