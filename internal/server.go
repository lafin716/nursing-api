//go:build wireinject
// +build wireinject

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	_ "nursing_api/docs"
	"nursing_api/internal/api"
	"nursing_api/internal/config"
	"nursing_api/internal/domain"
	"nursing_api/internal/middleware"
	"nursing_api/internal/router"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	"nursing_api/pkg/web"
)

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		router.Set,
		middleware.Set,
		api.Set,
		domain.Set,
	)))
}

type Server struct {
	app *fiber.App
	cfg *web.FiberConfig
	db  *ent.Client
}

func NewServer(
	cfg *web.FiberConfig,
	dbClient *database.DatabaseClient,
	router router.Routable,
) *Server {
	fiberClient := web.NewFiberClient(cfg)
	app := fiberClient.GetApp()
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	router.Init(&v1)

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
