package router

import (
	// _ "truphone-api/docs/app"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Truphone-api Swagger
// @version 1.0
// @description API Documentation

// @contact.name Luccas Machado
// @contact.url https://www.linkedin.com/in/luccas-machado-das-chagas-ab5897105/
// @contact.email luccaa.chagas23@gmail.com

// @host
// @BasePath /
func Handler() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(CorsConfig))
	e.GET("/health_check/", HealthCheck)
	e.POST("/auth", Auth)

	devices := e.Group("/devices")
	// devices.Use(middleware.JWTWithConfig(Config))

	devices.POST("/add", AddNewDevice)
	devices.GET("/all", GetAllDevices)
	devices.GET("/brand", GetDeviceByBrand)
	devices.GET("", GetDevicesByID)
	devices.DELETE("/remove", DeleteDeviceByID)
	devices.PATCH("/update", UpdateDevice) // Full or partial

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
