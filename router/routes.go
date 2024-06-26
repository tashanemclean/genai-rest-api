package router

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"

	"github.com/tashanemclean/genai-rest-api/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	// TODO config validatior for server

	// Health check route
	e.GET("/health", handlers.Healthcheck)
	e.GET("status", echoprometheus.NewHandler())
	e.GET("/v1", handlers.Default)

	v1 := e.Group("/v1")

	// TODO implement protect routes requiring authentication

	// Classification routes
	v1.POST("/classifytext", handlers.ClassifyText)
}