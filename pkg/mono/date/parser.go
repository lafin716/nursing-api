package date

import (
	"log"
	"time"
)

type parser struct {
}

type Parser interface {
	ParseToday(layout string) (time.Time, error)
	Parse(layout string, value string) (time.Time, error)
	ParseWithoutTime(layout string, value string) (time.Time, error)
	ParseWithTime(layout string, value string) (time.Time, error)
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
		return d.getToday(layout)
	}

	return t, nil
}

func (d parser) ParseWithoutTime(layout string, value string) (time.Time, error) {
	replaced := ReplaceLayout(layout)
	t, err := time.Parse(replaced, value)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func (d parser) ParseWithTime(layout string, value string) (time.Time, error) {
	// 현재일시에서 시간만 추출
	nowTime := time.Now().Format("15:04:05")
	log.Println("현재시간 : ", nowTime)
	replaced := ReplaceLayout(layout)
	t, err := time.Parse(replaced, value+" "+nowTime)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func (d parser) ParseForce(layout string, value string) time.Time {
	t, err := d.Parse(layout, value)
	if err != nil {
		t, _ = d.getToday(layout)
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

func (d parser) getToday(layout string) (time.Time, error) {
	t := time.Now()
	replaced := ReplaceLayout(layout)
	form := t.Format(replaced)

	return time.Parse(replaced, form)
}
