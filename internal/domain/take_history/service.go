package takehistory

import "nursing_api/pkg/mono"

type UseCase interface {
	GetList(req *GetListRequest) *GetListResponse
	GetDetail(req *GetDetailRequest) *GetDetailResponse
}

type takeHistoryService struct {
	takeHistory Repository
	mono        *mono.Client
}

func NewService(
	repo Repository,
	mono *mono.Client,
) UseCase {
	return &takeHistoryService{
		takeHistory: repo,
		mono:        mono,
	}
}

func (t takeHistoryService) GetList(req *GetListRequest) *GetListResponse {
	date, _ := t.mono.Date.Parse("Y-m-d", req.TargetDate)
	list, err := t.takeHistory.GetListByDate(req.UserId, date)
	if err != nil {
		return FailGetList("복용내역 목록 조회 중 오류가 발생하였습니다.", err)
	}

	for _, history := range list {
		items, err := t.takeHistory.GetItemsByHistoryId(req.UserId, history.ID, date)
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
