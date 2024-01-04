package container

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"nursing_api/app/utils"
	"nursing_api/pkg/auth"
	"nursing_api/pkg/database"
	"nursing_api/pkg/web"
	"os"
)

type FiberContainer struct {
	App       *fiber.App
	Ctx       *fiber.Ctx
	JwtClient *auth.JwtClient
	DBClient  *database.DatabaseClient
}

func NewFiberContainer(jwtClient *auth.JwtClient, dbClient *database.DatabaseClient) *FiberContainer {
	container := &FiberContainer{
		JwtClient: jwtClient,
		DBClient:  dbClient,
	}
	container.initContainer()

	return container
}

func (c *FiberContainer) initContainer() {
	if c.JwtClient == nil {
		log.Fatalln("JWT 모듈이 정의되지 않았습니다.")
	}

	if c.DBClient == nil {
		log.Fatalln("데이터베이스가 정의되지 않았습니다.")
	}

	appName := os.Getenv("APP_NAME")
	config := web.NewFiberConfig(appName)
	config.SetErrorHandler(globalErrorHandlerfunc)
	fiberClient := web.NewFiberClient(config)
	c.App = fiberClient.GetApp()
}

func globalErrorHandlerfunc(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return utils.ResponseEntity{
		Code:   code,
		Errors: err.Error(),
	}.ResponseError(c)
}
