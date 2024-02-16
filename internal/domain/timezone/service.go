package timezone

import (
	"github.com/google/uuid"
	"strings"
)

type timeZoneService struct {
	repo Repository
}

func NewTimeZoneService(
	repo Repository,
) UseCase {
	return &timeZoneService{
		repo: repo,
	}
}

func (t timeZoneService) GetList(userId uuid.UUID) ([]*TimeZone, error) {
	timezones, err := t.repo.GetTimeZones(userId)
	if err != nil {
		return nil, err
	}

	return timezones, nil
}

func (t timeZoneService) Create(req *CreateTimeZoneRequest) *CreateTimeZoneResponse {
	duplicate, _ := t.repo.GetDuplicate(
		req.UserId,
		req.Name,
		req.Midday,
		req.Hour,
		req.Minute,
	)

	if duplicate != nil {
		duplMsg := "중복되는 시간대가 존재합니다."
		if duplicate.Name == req.Name {
			duplMsg = "이미 존재하는 이름입니다."
		} else if duplicate.Hour == req.Hour && duplicate.Minute == req.Minute {
			duplMsg = "이미 등록된 시간입니다."
		}
		return FailCreateTimeZone(duplMsg, nil)
	}

	newTimeZone := &TimeZone{
		UserID: req.UserId,
		Name:   req.Name,
		Midday: req.Midday,
		Hour:   req.Hour,
		Minute: req.Minute,
	}

	timezone, err := t.repo.CreateTimeZone(newTimeZone)
	if err != nil {
		return FailCreateTimeZone("시간대 생성 중 오류가 발생하였습니다.", err)
	}

	return OkCreateTimeZone(timezone)
}

func (t timeZoneService) Update(req *UpdateTimeZoneRequest) *UpdateTimeZoneResponse {
	timezone, err := t.repo.GetTimeZone(req.ID, req.UserId)
	if err != nil {
		return FailUpdateTimeZone("시간대 정보를 찾을 수 없습니다.", err)
	}

	duplicate, _ := t.repo.GetDuplicate(
		req.UserId,
		req.Name,
		req.Midday,
		req.Hour,
		req.Minute,
	)

	if duplicate != nil && timezone.ID != duplicate.ID {
		duplMsg := "중복되는 시간대가 존재합니다."
		if duplicate.Name == req.Name {
			duplMsg = "이미 존재하는 이름입니다."
		} else if duplicate.Hour == req.Hour && duplicate.Minute == req.Minute {
			duplMsg = "이미 등록된 시간입니다."
		}
		return FailUpdateTimeZone(duplMsg, nil)
	}

	// 값이 있는 경우에만 도메인 값을 변경한다.
	if strings.TrimSpace(req.Name) != "" {
		timezone.Name = req.Name
	}
	if timezone.Midday != req.Midday {
		timezone.Midday = req.Midday
	}
	if timezone.Hour != req.Hour {
		timezone.Hour = req.Hour
	}
	if timezone.Minute != req.Minute {
		timezone.Minute = req.Minute
	}

	// 업데이트 처리
	updated, err := t.repo.UpdateTimeZone(timezone)
	if !updated || err != nil {
		return FailUpdateTimeZone("시간대 정보를 업데이트하는 중 오류가 발생하였습니다.", err)
	}

	return OkUpdateTimeZone(timezone)
}

func (t timeZoneService) Delete(req *DeleteTimeZoneRequest) *DeleteTimeZoneResponse {
	_, err := t.repo.GetTimeZone(req.ID, req.UserId)
	if err != nil {
		return FailDeleteTimeZone("시간대 정보를 찾을 수 없습니다.", err)
	}

	deleted, err := t.repo.DeleteTimeZone(req.ID, req.UserId)
	if deleted || err != nil {
		return FailDeleteTimeZone("시간대 삭제 중 오류가 발생하였습니다", err)
	}

	return OkDeleteTimeZone()
}
