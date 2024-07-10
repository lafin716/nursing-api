package date

import "time"

type formatter struct {
}

type Formatter interface {
	Format(t time.Time, layout string) string
	TruncateToDate(dateTime time.Time) time.Time
	TruncateToMonth(dateTime time.Time) time.Time
	TruncateToDateAddDay(dateTime time.Time, day int) time.Time
	TruncateToMonthAddMonth(dateTime time.Time, month int) time.Time
	ToFirstDayOfMonth(year, month int) time.Time
	GetDate(dateTime time.Time) string
	GetTime(dateTime time.Time) string
}

func NewFormatter() Formatter {
	return &formatter{}
}

func (f formatter) GetDate(dateTime time.Time) string {
	return dateTime.Format("2006-01-02")
}

func (f formatter) GetTime(dateTime time.Time) string {
	return dateTime.Format("15:04:05")
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

func (f formatter) TruncateToMonth(dateTime time.Time) time.Time {
	return time.Date(
		dateTime.Year(),
		dateTime.Month(),
		1,
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

func (f formatter) TruncateToMonthAddMonth(dateTime time.Time, month int) time.Time {
	return time.Date(
		dateTime.Year(),
		dateTime.Month()+time.Month(month),
		1,
		0,
		0,
		0,
		0,
		dateTime.Location(),
	)
}

func (f formatter) ToFirstDayOfMonth(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
}
