package date

type TimeLayout string

const (
	YEAR   = TimeLayout("2006")
	MONTH  = TimeLayout("01")
	DATE   = TimeLayout("02")
	HOUR   = TimeLayout("15")
	MINUTE = TimeLayout("04")
	SECOND = TimeLayout("05")
)
