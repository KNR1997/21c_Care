package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-echo-starter/docs"
	"go-echo-starter/internal/config"
	"go-echo-starter/internal/db"
	"go-echo-starter/internal/repositories"
	"go-echo-starter/internal/server"
	"go-echo-starter/internal/server/handlers"
	"go-echo-starter/internal/server/middleware"
	"go-echo-starter/internal/server/routes"
	"go-echo-starter/internal/services/ai"
	"go-echo-starter/internal/services/auth"
	"go-echo-starter/internal/services/billing"
	"go-echo-starter/internal/services/clinicalnote"
	"go-echo-starter/internal/services/drugcatalog"
	"go-echo-starter/internal/services/labtest"
	"go-echo-starter/internal/services/labtestcatalog"
	"go-echo-starter/internal/services/oauth"
	"go-echo-starter/internal/services/patient"
	"go-echo-starter/internal/services/post"
	"go-echo-starter/internal/services/prescribeddrug"
	"go-echo-starter/internal/services/report"
	"go-echo-starter/internal/services/settings"
	"go-echo-starter/internal/services/token"
	"go-echo-starter/internal/services/user"
	"go-echo-starter/internal/services/visit"
	"go-echo-starter/internal/slogx"

	"github.com/caarlos0/env/v11"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

const shutdownTimeout = 20 * time.Second

//	@title			ABC Clinic
//	@version		1.0
//	@description	This is a demo version of Echo app.

//	@contact.name	KNR Solutions
//	@contact.email	kethakaranasinghe@gmail.com

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/
func main() {
	if err := run(); err != nil {
		slog.Error("Service run error", "err", err.Error())
		os.Exit(1)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("load env file: %w", err)
	}

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("parse env: %w", err)
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	if err := slogx.Init(cfg.Logger); err != nil {
		return fmt.Errorf("init logger: %w", err)
	}

	gormDB, err := db.NewGormDB(cfg.DB)
	if err != nil {
		return fmt.Errorf("new db connection: %w", err)
	}

	userRepository := repositories.NewUserRepository(gormDB)
	userService := user.NewService(userRepository)

	postRepository := repositories.NewPostRepository(gormDB)
	postService := post.NewService(postRepository)

	patientRepository := repositories.NewPatientRepository(gormDB)
	patientService := patient.NewService(patientRepository)

	drugCatalogRepository := repositories.NewDrugCatalogRepository(gormDB)
	drugCatalogService := drugcatalog.NewService(drugCatalogRepository)

	labTestcatalogRepository := repositories.NewLabTestCatalogRepository(gormDB)
	labTestCatalogService := labtestcatalog.NewService(labTestcatalogRepository)

	labTestRepository := repositories.NewLabTestRepository(gormDB)
	labTestService := labtest.NewService(labTestRepository)

	clinicalNoteRepository := repositories.NewClinicalNoteRepository(gormDB)
	clinicalNoteService := clinicalnote.NewService(clinicalNoteRepository)

	prescribedDrugRepository := repositories.NewPrescribedDrugRepository(gormDB)
	prescribedDrugService := prescribeddrug.NewService(prescribedDrugRepository)

	aiService := ai.NewService(cfg.APIKey)

	settingsService := settings.NewService()

	visitRepository := repositories.NewVisitRepository(gormDB)
	visitService := visit.NewService(
		visitRepository,
		aiService,
		labTestService,
		clinicalNoteService,
		drugCatalogRepository,
		labTestcatalogRepository,
		prescribedDrugService,
	)

	billingService := billing.NewService(
		repositories.NewBillingRepository(gormDB),
		prescribedDrugRepository,
		labTestRepository,
	)

	reportService := report.NewService(
		visitRepository,
	)

	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return fmt.Errorf("oidc.NewProvider: %w", err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.OAuth.ClientID})

	tokenService := token.NewService(
		time.Now,
		cfg.Auth.AccessTokenDuration,
		cfg.Auth.RefreshTokenDuration,
		[]byte(cfg.Auth.AccessSecret),
		[]byte(cfg.Auth.RefreshSecret),
	)

	authService := auth.NewService(userService, tokenService)
	oAuthService := oauth.NewService(verifier, tokenService, userService)

	settingsHandler := handlers.NewSettingsHandler(settingsService)
	drugCatalogHandler := handlers.NewDrugCatalogHandlers(drugCatalogService)
	labTestCatalogHandler := handlers.NewLabTestCatalogHandlers(labTestCatalogService)
	visitHandler := handlers.NewVisitHandlers(visitService)
	billingHandler := handlers.NewBillingHandlers(billingService)
	reportHandler := handlers.NewReportHandler(reportService)
	patientHandler := handlers.NewPatientHandler(patientService)
	postHandler := handlers.NewPostHandlers(postService)
	authHandler := handlers.NewAuthHandler(authService)
	oAuthHandler := handlers.NewOAuthHandler(oAuthService)
	registerHandler := handlers.NewRegisterHandler(userService)

	authMiddleware := middleware.NewAuthMiddleware(cfg.Auth.AccessSecret)
	reguestLoggerMiddleware := middleware.NewRequestLogger(slogx.NewTraceStarter(uuid.NewV7))
	requestDebuggerMiddleware := middleware.NewRequestDebugger()

	engine := routes.ConfigureRoutes(routes.Handlers{
		SettingsHandler:           settingsHandler,
		DrugCatalogHandler:        drugCatalogHandler,
		LabTestCatalogHandler:     labTestCatalogHandler,
		VisitHandler:              visitHandler,
		BillingHandler:            billingHandler,
		ReportHandler:             reportHandler,
		PatientHandler:            patientHandler,
		PostHandler:               postHandler,
		AuthHandler:               authHandler,
		OAuthHandler:              oAuthHandler,
		RegisterHandler:           registerHandler,
		AuthMiddleware:            authMiddleware,
		RequestLoggerMiddleware:   reguestLoggerMiddleware,
		RequestDebuggerMiddleware: requestDebuggerMiddleware,
	})
	if err != nil {
		return fmt.Errorf("configure routes: %w", err)
	}

	app := server.NewServer(engine)
	go func() {
		if err = app.Start(cfg.HTTP.Port); err != nil {
			slog.Error("Server error", "err", err.Error())
		}
	}()

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	<-shutdownChannel

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := app.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("http server shutdown: %w", err)
	}

	dbConnection, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("get db connection: %w", err)
	}

	if err := dbConnection.Close(); err != nil {
		return fmt.Errorf("close db connection: %w", err)
	}

	return nil
}
