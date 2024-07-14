package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/timezone"
	"nursing_api/pkg/jwt"
)

type TimeZoneApi interface {
	GetList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	IsDeletable(c *fiber.Ctx) error
}

type timezoneApi struct {
	service   timezone.UseCase
	jwtClient *jwt.JwtClient
}

func NewTimeZoneApi(
	service timezone.UseCase,
	jwtClient *jwt.JwtClient,
) TimeZoneApi {
	return &timezoneApi{
		service:   service,
		jwtClient: jwtClient,
	}
}

// @summary 복용시간대 목록 조회
// @description 복용시간대 목록을 조회하는 엔드포인트
// @produce json
// @router /timezone [get]
// @Security Bearer
func (t timezoneApi) GetList(c *fiber.Ctx) error {
	userId, err := getUserId(t.jwtClient, c)
	if err != nil {
		return err
	}

	resp := t.service.GetList(userId)
	return ResolveResponse(resp, c)
}

// @summary 복용시간대 추가
// @description 복용시간대 템플릿을 생성하는 엔드포인트
// @accept json
// @produce json
// @param dto body timezone.CreateTimeZoneRequest true "복용시간대 정보"
// @router /timezone [post]
// @Security Bearer
func (t timezoneApi) Create(c *fiber.Ctx) error {
	userId, err := getUserId(t.jwtClient, c)
	if err != nil {
		return err
	}

	req := new(timezone.CreateTimeZoneRequest)
	err = c.BodyParser(req)
	if err != nil {
		return FailParam(err.Error(), c)
	}

	errs := validateParameter(req)
	if errs != nil {
		return FailParam(errs, c)
	}
	req.UserId = userId

	resp := t.service.Create(req)
	return ResolveResponse(resp, c)
}

// @summary 복용시간대 업데이트
// @description 복용시간대 정보를 업데이트하는 엔드포인트
// @accept json
// @produce json
// @router /timezone [patch]
// @Security Bearer
func (t timezoneApi) Update(c *fiber.Ctx) error {
	userId, err := getUserId(t.jwtClient, c)
	if err != nil {
		return err
	}

	req := new(timezone.UpdateTimeZoneRequest)
	err = c.BodyParser(req)
	if err != nil {
		return FailParam(err.Error(), c)
	}
	req.UserId = userId

	resp := t.service.Update(req)
	return ResolveResponse(resp, c)
}

// @summary 복용시간대 삭제
// @description 복용시간대를 삭제하는 엔드포인트
// @accept json
// @produce json
// @param id path string true "복용시간대 고유번호"
// @router /timezone/:id [delete]
// @Security Bearer
func (t timezoneApi) Delete(c *fiber.Ctx) error {
	userId, err := getUserId(t.jwtClient, c)
	if err != nil {
		return err
	}

	req := new(timezone.DeleteTimeZoneRequest)
	err = c.ParamsParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err.Error()).
			Error(c)
	}

	errs := validateParameter(req)
	if errs != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(errs).
			Error(c)
	}
	req.UserId = userId

	resp := t.service.Delete(req)
	return ResolveResponse(resp, c)
}

// @summary 복용시간대 삭제 가능 여부
// @description 복용시간대를 삭제 가능 여부를 확인하는 엔드포인트
// @produce json
// @param id path string true "복용시간대 고유번호"
// @router /timezone/deletable/:id [get]
// @Security Bearer
func (t timezoneApi) IsDeletable(c *fiber.Ctx) error {
	userId, err := getUserId(t.jwtClient, c)
	if err != nil {
		return err
	}

	req := new(timezone.IsDeletableTimeZoneRequest)
	err = c.ParamsParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err.Error()).
			Error(c)
	}

	errs := validateParameter(req)
	if errs != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(errs).
			Error(c)
	}
	req.UserId = userId

	resp := t.service.IsDeletable(req)
	return ResolveResponse(resp, c)
}
