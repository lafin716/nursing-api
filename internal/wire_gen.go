// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"nursing_api/internal/api"
	"nursing_api/internal/config"
	"nursing_api/internal/domain/auth"
	"nursing_api/internal/domain/medicine"
	"nursing_api/internal/domain/plan"
	"nursing_api/internal/domain/prescription"
	"nursing_api/internal/domain/take_history"
	"nursing_api/internal/domain/timezone"
	"nursing_api/internal/domain/timezone_link"
	"nursing_api/internal/domain/user"
	"nursing_api/internal/middleware"
	"nursing_api/internal/router"
	"nursing_api/pkg/api/medicine"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	"nursing_api/pkg/jwt"
	"nursing_api/pkg/mono"
	"nursing_api/pkg/web"
)

import (
	_ "nursing_api/docs"
)

// Injectors from server.go:

func New() (*Server, error) {
	fiberConfig := config.NewFiberConfig()
	databaseConfig := config.NewDatabaseConfig()
	databaseClient := database.NewPostgresClient(databaseConfig)
	repository := auth.NewRepository(databaseClient)
	jwtConfig := config.NewJwtConfig()
	jwtClient := jwt.NewJwtClient(jwtConfig)
	tokenVerifyMiddleware := middleware.NewJwtMiddleware(repository, jwtClient)
	mainHttpApi := api.NewMainHttpApi()
	userRepository := user.NewRepository(databaseClient)
	useCase := user.NewService(userRepository)
	authUseCase := auth.NewService(useCase, repository, jwtClient)
	authHttpApi := api.NewAuthHttpApi(authUseCase, jwtClient)
	userHttpApi := api.NewUserHttpApi(useCase, jwtClient)
	medicineRepository := medicine.NewRepository(databaseClient)
	medicineApiConfig := config.NewMedicineApiConfig()
	medicineApi := medicine_api.NewMedicineApi(medicineApiConfig)
	medicineUseCase := medicine.NewService(medicineRepository, medicineApi)
	medicineHttpApi := api.NewMedicineHttpApi(medicineUseCase)
	prescriptionRepository := prescription.NewRepository(databaseClient)
	prescriptionUseCase := prescription.NewService(prescriptionRepository, jwtClient)
	prescriptionApi := api.NewPrescriptionApi(prescriptionUseCase, jwtClient)
	takehistoryRepository := takehistory.NewRepository(databaseClient)
	takehistoryUseCase := takehistory.NewService(takehistoryRepository, prescriptionRepository)
	takeHistoryHttpApi := api.NewTakeHistoryHttpApi(takehistoryUseCase, jwtClient)
	client := mono.NewMono()
	timezoneRepository := timezone.NewRepository(databaseClient)
	timezonelinkRepository := timezonelink.NewRepository(databaseClient)
	planUseCase := plan.NewService(databaseClient, client, prescriptionRepository, timezoneRepository, timezonelinkRepository, takehistoryRepository, medicineRepository)
	planHttpApi := api.NewPlanHttpApi(planUseCase, jwtClient)
	timezoneUseCase := timezone.NewService(timezoneRepository)
	timeZoneApi := api.NewTimeZoneApi(timezoneUseCase, jwtClient)
	routable := router.NewRouter(tokenVerifyMiddleware, mainHttpApi, authHttpApi, userHttpApi, medicineHttpApi, prescriptionApi, takeHistoryHttpApi, planHttpApi, timeZoneApi)
	server := NewServer(fiberConfig, databaseClient, routable)
	return server, nil
}

// server.go:

type Server struct {
	app *fiber.App
	cfg *web.FiberConfig
	db  *ent.Client
}

func NewServer(
	cfg *web.FiberConfig,
	dbClient *database.DatabaseClient, router2 router.Routable,

) *Server {
	fiberClient := web.NewFiberClient(cfg)
	app := fiberClient.GetApp()
	app.Get("/swagger/*", swagger.HandlerDefault)
	api2 := app.Group("/api")
	v1 := api2.Group("/v1")
	router2.
		Init(&v1)

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
