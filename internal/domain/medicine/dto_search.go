package medicine

type SearchPillResponse struct {
	Success bool
	Message string
	Pills   []Medicine
	Error   error
}

func OkSearchPill(pills []Medicine) *SearchPillResponse {
	return &SearchPillResponse{
		Success: true,
		Message: "의약품 검색 완료",
		Pills:   pills,
	}
}

func FailSearchPill(message string, err error) *SearchPillResponse {
	return &SearchPillResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
