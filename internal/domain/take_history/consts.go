package takehistory

type TakeStatus string
type TakePillStatus string

const (
	DONE    = TakeStatus("DONE")
	PARTIAL = TakeStatus("PARTIAL")
	NEVER   = TakeStatus("NEVER")
)

const (
	Y = TakePillStatus("Y")
	N = TakePillStatus("N")
)
