package routes

import (
	"net/http"

	"go-echo-starter/internal/server/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	DrugCatalogHandler    *handlers.DrugCatalogHandlers
	LabTestCatalogHandler *handlers.LabTestCatalogHandlers
	PatientHandler        *handlers.PatientHandlers
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

	privateAPI.POST("/login", handlers.AuthHandler.Login)
	privateAPI.POST("/register", handlers.RegisterHandler.Register)

	// Authorized API route initialization.
	//
	// These endpoints implement the core application logic and require authentication
	// before they can be accessed.
	authorizedAPI := api.Group("", handlers.RequestDebuggerMiddleware, handlers.AuthMiddleware)

	authorizedAPI.POST("/posts", handlers.PostHandler.CreatePost)
	authorizedAPI.GET("/posts", handlers.PostHandler.GetPosts)
	authorizedAPI.PUT("/posts/:id", handlers.PostHandler.UpdatePost)
	authorizedAPI.DELETE("/posts/:id", handlers.PostHandler.DeletePost)

	authorizedAPI.POST("/drugCatalogs", handlers.DrugCatalogHandler.CreateDrugCatalog)
	authorizedAPI.GET("/drugCatalogs", handlers.DrugCatalogHandler.GetDrugCatalogs)
	authorizedAPI.PUT("/drugCatalogs/:id", handlers.DrugCatalogHandler.UpdateDrugCatalog)
	authorizedAPI.DELETE("/drugCatalogs/:id", handlers.DrugCatalogHandler.DeleteDrugCatalog)

	authorizedAPI.POST("/labTestCatalogs", handlers.LabTestCatalogHandler.CreateLabTestCatalog)
	authorizedAPI.GET("/labTestCatalogs", handlers.LabTestCatalogHandler.GetLabTestCatalogs)
	authorizedAPI.PUT("/labTestCatalogs/:id", handlers.LabTestCatalogHandler.UpdateLabTestCatalog)
	authorizedAPI.DELETE("/labTestCatalogs/:id", handlers.LabTestCatalogHandler.DeleteLabTestCatalog)

	authorizedAPI.POST("/patients", handlers.PatientHandler.CreatePatient)
	authorizedAPI.GET("/patients", handlers.PatientHandler.GetPatients)
	authorizedAPI.PUT("/patients/:id", handlers.PatientHandler.UpdatePatient)
	authorizedAPI.DELETE("/patients/:id", handlers.PatientHandler.DeletePatient)

	return engine

}
