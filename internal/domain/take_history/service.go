package takehistory

type takeHistoryService struct {
	repo TakeHistoryRepository
}

func NewTakeHistoryService(
	repo TakeHistoryRepository,
) TakeHistoryUseCase {
	return &takeHistoryService{
		repo: repo,
	}
}

func (t takeHistoryService) GetList(req *GetListRequest) *GetListResponse {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryService) GetDetail(req *GetDetailRequest) *GetDetailResponse {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryService) TakePlanToggle(req *TakePlanRequest) *TakePlanResponse {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryService) TakePillToggle(req *TakePillRequest) *TakePillResponse {
	//TODO implement me
	panic("implement me")
}
