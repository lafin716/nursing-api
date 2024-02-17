package date

import "time"

type formatter struct {
}

type Formatter interface {
	Format(t time.Time, layout string) string
}

func NewFormatter() Formatter {
	return &formatter{}
}

func (f formatter) Format(t time.Time, layout string) string {
	replaced := ReplaceLayout(layout)
	return t.Format(replaced)
}
