//go:build wireinject
// +build wireinject

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"nursing_api/internal/api"
	"nursing_api/internal/config"
	"nursing_api/internal/infrastructure/persistence"
	"nursing_api/internal/router"
	"nursing_api/internal/service"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	"nursing_api/pkg/jwt"
	"nursing_api/pkg/web"
)

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		router.Set,
		api.Set,
		service.Set,
		persistence.Set,
		jwt.Set,
		database.Set,
	)))
}

type Server struct {
	app *fiber.App
	cfg *web.FiberConfig
	db  *ent.Client
}

func NewServer(
	cfg *web.FiberConfig,
	jwtClient *jwt.JwtClient,
	dbClient *database.DatabaseClient,
	authRouter router.AuthRouter,
	userRouter router.UserRouter,
) *Server {
	fiberClient := web.NewFiberClient(cfg)
	app := fiberClient.GetApp()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	authRouter.Init(&v1, jwtClient.Middleware)
	userRouter.Init(&v1, jwtClient.Middleware)

	return &Server{
		app: app,
		cfg: cfg,
		db:  dbClient.Client,
	}
}

func (serv *Server) App() *fiber.App {
	return serv.app
}

func (serv *Server) Config() *web.FiberConfig {
	return serv.cfg
}

func (serv *Server) DB() *ent.Client {
	return serv.db
}
