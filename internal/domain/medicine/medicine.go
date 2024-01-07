package medicine

import "time"

type Medicine struct {
	Name           string
	EffectMemo     string
	SideEffectMemo string
	Warning        string
	Caution        string
	Usage          string
	Keeping        string
	CreatedAt      time.Time
}
