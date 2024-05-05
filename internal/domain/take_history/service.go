package takehistory

import (
	"nursing_api/internal/common/dto"
	"nursing_api/internal/common/response"
	"nursing_api/pkg/mono"
)

type UseCase interface {
	GetList(req *GetListRequest) dto.BaseResponse[[]*TakeHistory]
	GetDetail(req *GetDetailRequest) dto.BaseResponse[TakeHistory]
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

func (t takeHistoryService) GetList(req *GetListRequest) dto.BaseResponse[[]*TakeHistory] {
	date, _ := t.mono.Date.Parse("Y-m-d", req.TargetDate)
	list, err := t.takeHistory.GetListByDate(req.UserId, date)
	if err != nil {
		return dto.Fail[[]*TakeHistory](response.CODE_FAIL_GET_LIST_TAKE_HISTORY, err)
	}

	for _, history := range list {
		items, err := t.takeHistory.GetItemsByHistoryId(req.UserId, history.ID, date)
		if err != nil {
			return dto.Fail[[]*TakeHistory](response.CODE_FAIL_GET_ITEM_LIST_TAKE_HISTORY, err)
		}

		history.TakeHistoryItems = items
	}

	return dto.Ok[[]*TakeHistory](response.CODE_SUCCESS, &list)
}

func (t takeHistoryService) GetDetail(req *GetDetailRequest) dto.BaseResponse[TakeHistory] {
	//TODO implement me
	panic("implement me")
}
