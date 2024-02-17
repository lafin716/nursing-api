package date

import "time"

type parser struct {
}

type Parser interface {
	Parse(layout string, value string) (time.Time, error)
	ParseForce(layout string, value string) time.Time
	ParseWithDefault(layout string, value string, zero time.Time) time.Time
}

func NewParser() Parser {
	return &parser{}
}

func (d parser) ParseToday(layout string) (time.Time, error) {
	replaced := ReplaceLayout(layout)
	value := time.Now().Format(replaced)
	t, err := time.Parse(replaced, value)
	if err != nil {
		return time.Now(), err
	}

	return t, nil
}

func (d parser) Parse(layout string, value string) (time.Time, error) {
	replaced := ReplaceLayout(layout)
	t, err := time.Parse(replaced, value)
	if err != nil {
		return time.Now(), err
	}

	return t, nil
}

func (d parser) ParseForce(layout string, value string) time.Time {
	t, err := d.Parse(layout, value)
	if err != nil {
		return time.Now()
	}

	return t
}

func (d parser) ParseWithDefault(layout string, value string, zero time.Time) time.Time {
	t, err := d.Parse(layout, value)
	if err != nil {
		return zero
	}

	return t
}
