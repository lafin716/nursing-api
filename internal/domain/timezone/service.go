package timezone

import (
	"github.com/google/uuid"
	"nursing_api/internal/common/dto"
	"nursing_api/internal/common/response"
	"strings"
)

type UseCase interface {
	GetList(userId uuid.UUID) dto.BaseResponse[[]*TimeZone]
	Create(req *CreateTimeZoneRequest) dto.BaseResponse[TimeZone]
	Update(req *UpdateTimeZoneRequest) dto.BaseResponse[TimeZone]
	Delete(req *DeleteTimeZoneRequest) dto.BaseResponse[bool]
}

type service struct {
	repo Repository
}

func NewService(
	repo Repository,
) UseCase {
	return &service{
		repo: repo,
	}
}

func (t service) GetList(userId uuid.UUID) dto.BaseResponse[[]*TimeZone] {
	timezones, err := t.repo.GetTimeZones(userId)
	if err != nil {
		return dto.Fail[[]*TimeZone](response.CODE_FAIL_GET_LIST_TAKE_TIMEZONE, err)
	}

	return dto.Ok[[]*TimeZone](response.CODE_SUCCESS, &timezones)
}

func (t service) Create(req *CreateTimeZoneRequest) dto.BaseResponse[TimeZone] {
	duplicate, _ := t.repo.GetDuplicate(
		req.UserId,
		req.Name,
		req.Midday,
		req.Hour,
		req.Minute,
	)

	if duplicate != nil {
		errCode := response.CODE_ALREADY_EXIST_TIMEZONE_NAME
		if duplicate.Name == req.Name {
			errCode = response.CODE_ALREADY_EXIST_TIMEZONE_NAME
		} else if duplicate.Hour == req.Hour && duplicate.Minute == req.Minute {
			errCode = response.CODE_ALREADY_EXIST_TIMEZONE_HOUR_MINUTE
		}
		return dto.Fail[TimeZone](errCode, nil)
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
		return dto.Fail[TimeZone](response.CODE_FAIL_DURING_CREATE_TIMEZONE, err)
	}

	return dto.Ok[TimeZone](response.CODE_SUCCESS, timezone)
}

func (t service) Update(req *UpdateTimeZoneRequest) dto.BaseResponse[TimeZone] {
	timezone, err := t.repo.GetTimeZone(req.ID, req.UserId)
	if err != nil {
		return dto.Fail[TimeZone](response.CODE_NOT_FOUND_TIMEZONE, err)
	}

	duplicate, _ := t.repo.GetDuplicate(
		req.UserId,
		req.Name,
		req.Midday,
		req.Hour,
		req.Minute,
	)

	if duplicate != nil && timezone.ID != duplicate.ID {
		errCode := response.CODE_ALREADY_EXIST_TIMEZONE_NAME
		if duplicate.Name == req.Name {
			errCode = response.CODE_ALREADY_EXIST_TIMEZONE_NAME
		} else if duplicate.Hour == req.Hour && duplicate.Minute == req.Minute {
			errCode = response.CODE_ALREADY_EXIST_TIMEZONE_HOUR_MINUTE
		}
		return dto.Fail[TimeZone](errCode, nil)
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
		return dto.Fail[TimeZone](response.CODE_FAIL_DURING_UPDATE_TIMEZONE, err)
	}

	return dto.Ok[TimeZone](response.CODE_SUCCESS, timezone)
}

func (t service) Delete(req *DeleteTimeZoneRequest) dto.BaseResponse[bool] {
	_, err := t.repo.GetTimeZone(req.ID, req.UserId)
	if err != nil {
		return dto.Fail[bool](response.CODE_NOT_FOUND_TIMEZONE, err)
	}

	deleted, err := t.repo.DeleteTimeZone(req.ID, req.UserId)
	if err != nil {
		return dto.Fail[bool](response.CODE_FAIL_DURING_DELETE_TIMEZONE, err)
	}

	return dto.Ok[bool](response.CODE_SUCCESS, &deleted)
}
