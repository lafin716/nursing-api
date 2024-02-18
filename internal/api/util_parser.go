package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"nursing_api/pkg/jwt"
)

type ParamType int

const (
	QUERY = ParamType(0)
	PARAM = ParamType(1)
	BODY  = ParamType(2)
)

type paramParser struct {
	paramType ParamType
	dto       interface{}
	userId    uuid.UUID
	jwtClient *jwt.JwtClient
	ctx       *fiber.Ctx
	err       error
}

type ParamParser interface {
	init()
	Error() error
	GetUserId() uuid.UUID
}

func ParseRequest(dto interface{}, paramType ParamType, jwtClient *jwt.JwtClient, c *fiber.Ctx) ParamParser {
	parser := &paramParser{
		paramType: paramType,
		dto:       dto,
		jwtClient: jwtClient,
		ctx:       c,
	}
	parser.init()

	return parser
}

func (p paramParser) init() {
	err := p.parse()
	if err != nil {
		p.err = err
		return
	}

	errs := validateParameter(p.dto)
	if errs != nil {
		p.err = FailParam(errs, p.ctx)
		return
	}

	userId, err := getUserId(p.jwtClient, p.ctx)
	if err != nil {
		p.err = FailAuth(err.Error(), p.ctx)
		return
	}

	p.userId = userId
}

func (p paramParser) parse() error {
	var err error
	switch p.paramType {
	case QUERY:
		err = p.ctx.QueryParser(p.dto)
	case PARAM:
		err = p.ctx.ParamsParser(p.dto)
	case BODY:
		err = p.ctx.BodyParser(p.dto)
	}

	return err
}

func (p paramParser) GetUserId() uuid.UUID {
	return p.userId
}

func (p paramParser) Error() error {
	return p.err
}
