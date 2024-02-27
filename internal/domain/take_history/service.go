package takehistory

type UseCase interface {
	GetList(req *GetListRequest) *GetListResponse
	GetDetail(req *GetDetailRequest) *GetDetailResponse
}

type takeHistoryService struct {
	takeHistory Repository
}

func NewService(
	repo Repository,
) UseCase {
	return &takeHistoryService{
		takeHistory: repo,
	}
}

func (t takeHistoryService) GetList(req *GetListRequest) *GetListResponse {
	list, err := t.takeHistory.GetListByDate(req.UserId, req.TargetDate)
	if err != nil {
		return FailGetList("복용내역 목록 조회 중 오류가 발생하였습니다.", err)
	}

	for _, history := range list {
		items, err := t.takeHistory.GetItemsByHistoryId(req.UserId, history.ID, req.TargetDate)
		if err != nil {
			return FailGetList("복용내역 아이템 목록 조회 중 오류가 발생하였습니다.", err)
		}

		history.TakeHistoryItems = items
	}

	return OkGetList(list)
}

func (t takeHistoryService) GetDetail(req *GetDetailRequest) *GetDetailResponse {
	//TODO implement me
	panic("implement me")
}
