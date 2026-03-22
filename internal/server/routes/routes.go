package routes

import (
	"net/http"

	"go-echo-starter/internal/server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	SettingsHandler       *handlers.SettingsHandler
	DrugCatalogHandler    *handlers.DrugCatalogHandler
	LabTestCatalogHandler *handlers.LabTestCatalogHandler
	PatientHandler        *handlers.PatientHandler
	VisitHandler          *handlers.VisitHandlers
	BillingHandler        *handlers.BillingHandler
	ReportHandler         *handlers.ReportHandler
	PostHandler           *handlers.PostHandlers
	AuthHandler           *handlers.AuthHandler
	OAuthHandler          *handlers.OAuthHandler
	RegisterHandler       *handlers.RegisterHandler

	AuthMiddleware            echo.MiddlewareFunc
	RequestLoggerMiddleware   echo.MiddlewareFunc
	RequestDebuggerMiddleware echo.MiddlewareFunc
}

func ConfigureRoutes(handlers Handlers) *echo.Echo {
	engine := echo.New()

	// Enable CORS
	engine.Use(middleware.CORS())

	// Technical API route initialization.
	// This works with Echo v4
	engine.GET("/swagger/*", echoSwagger.WrapHandler)
	engine.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	api := engine.Group("", handlers.RequestLoggerMiddleware)
	// Private API routes initialization.
	//
	// These endpoints are used primarily for authentication/authorization and may carry sensitive data.
	// Do NOT log request or response bodies; doing so could expose client information.
	privateAPI := api.Group("")

	privateAPI.GET("/settings", handlers.SettingsHandler.Get)
	privateAPI.POST("/login", handlers.AuthHandler.Login)
	privateAPI.POST("/register", handlers.RegisterHandler.Register)
	privateAPI.POST("/logout", handlers.AuthHandler.Logout)

	// Authorized API route initialization.
	//
	// These endpoints implement the core application logic and require authentication
	// before they can be accessed.
	authorizedAPI := api.Group("", handlers.RequestDebuggerMiddleware, handlers.AuthMiddleware)

	authorizedAPI.GET("/me", handlers.AuthHandler.Me)

	authorizedAPI.POST("/posts", handlers.PostHandler.CreatePost)
	authorizedAPI.GET("/posts", handlers.PostHandler.GetPosts)
	authorizedAPI.PUT("/posts/:id", handlers.PostHandler.UpdatePost)
	authorizedAPI.DELETE("/posts/:id", handlers.PostHandler.DeletePost)

	authorizedAPI.POST("/drugs", handlers.DrugCatalogHandler.Create)
	authorizedAPI.GET("/drugs", handlers.DrugCatalogHandler.ListPaginated)
	authorizedAPI.PUT("/drugs/:id", handlers.DrugCatalogHandler.Update)
	authorizedAPI.GET("/drugs/:id", handlers.DrugCatalogHandler.Get)
	authorizedAPI.DELETE("/drugs/:id", handlers.DrugCatalogHandler.Delete)

	authorizedAPI.POST("/labTests", handlers.LabTestCatalogHandler.Create)
	authorizedAPI.GET("/labTests", handlers.LabTestCatalogHandler.ListPaginated)
	authorizedAPI.PUT("/labTests/:id", handlers.LabTestCatalogHandler.Update)
	authorizedAPI.GET("/labTests/:id", handlers.LabTestCatalogHandler.Get)
	authorizedAPI.DELETE("/labTests/:id", handlers.LabTestCatalogHandler.Delete)

	authorizedAPI.POST("/patients", handlers.PatientHandler.Create)
	authorizedAPI.GET("/patients", handlers.PatientHandler.ListPaginated)
	authorizedAPI.PUT("/patients/:id", handlers.PatientHandler.Update)
	authorizedAPI.GET("/patients/:id", handlers.PatientHandler.Get)
	authorizedAPI.DELETE("/patients/:id", handlers.PatientHandler.Delete)

	authorizedAPI.POST("/visits", handlers.VisitHandler.Create)
	authorizedAPI.GET("/visits", handlers.VisitHandler.ListPaginated)
	authorizedAPI.GET("/visits/:id", handlers.VisitHandler.Get)
	authorizedAPI.POST("/visits/preview", handlers.VisitHandler.Preview)
	authorizedAPI.GET("/visits/:id/report", handlers.ReportHandler.GeneratePDF)

	authorizedAPI.GET("/billings", handlers.BillingHandler.ListPaginated)
	authorizedAPI.GET("/billings/:id", handlers.BillingHandler.Get)
	authorizedAPI.POST("/billings", handlers.BillingHandler.Create)

	return engine
}
