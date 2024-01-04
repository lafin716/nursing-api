package bootstrap

import (
	"nursing_api/app/container"
	"nursing_api/app/controllers"
	"nursing_api/app/routes"
	"nursing_api/pkg/auth"
	"nursing_api/pkg/database"
)

type FiberApplication struct {
	container *container.FiberContainer
}

func NewFiberApplication(jwtClient *auth.JwtClient, dbClient *database.DatabaseClient) *FiberApplication {
	application := &FiberApplication{
		container: container.NewFiberContainer(jwtClient, dbClient),
	}

	application.initApp()
	return application
}

func (c *FiberApplication) initApp() {
	handler := controllers.NewFiberHandler(c.container)
	routes.NewRoutes(handler).Resolve()
}
