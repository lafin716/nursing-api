package plan

type AddPlanRequest struct {
	Name      string  `json:"name"`
	Hospital  string  `json:"hospital"`
	StartedAt string  `json:"started_at"`
	TakeDays  float64 `json:"take_days"`
	Memo      string  `json:"memo"`
	Items     string  `json:"items"`
}

type AddPlanItemsRequest struct {
}
