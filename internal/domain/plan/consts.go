package plan

const (
	TTZ_MORNING = "MORNING"
	TTZ_LUNCH   = "LUNCH"
	TTZ_DINNER  = "DINNER"
	TTZ_SLEEP   = "SLEEP"
	TTZ_ETC     = "ETC"
)

const (
	TM_BEFORE = "BEFORE"
	TM_AFTER  = "AFTER"
	TM_ANY    = "ANY"
)

const (
	TXT_TTZ_MORNING = "아침"
	TXT_TTZ_LUNCH   = "점심"
	TXT_TTZ_DINNER  = "저녁"
	TXT_TTZ_SLEEP   = "취침전"
	TXT_TTZ_ETC     = "기타"
)

const (
	TXT_TM_BEFORE = "식전"
	TXT_TM_AFTER  = "식후"
	TXT_TM_ANY    = "수시"
)

var timezoneText = map[string]string{
	TTZ_MORNING: TXT_TTZ_MORNING,
	TTZ_LUNCH:   TXT_TTZ_LUNCH,
	TTZ_DINNER:  TXT_TTZ_DINNER,
	TTZ_SLEEP:   TXT_TTZ_SLEEP,
	TTZ_ETC:     TXT_TTZ_ETC,
}

var takeTimingText = map[string]string{
	TM_BEFORE: TXT_TM_BEFORE,
	TM_AFTER:  TXT_TM_AFTER,
	TM_ANY:    TXT_TM_ANY,
}

func GetTimeZone(timeZone string) string {
	return timezoneText[timeZone]
}

func GetTakeTiming(timing string) string {
	return takeTimingText[timing]
}
