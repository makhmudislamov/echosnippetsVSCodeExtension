package packagename

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handlers
func create(c echo.Context) error {
	return nil
}

func find(c echo.Context) error {
	return nil
}

func update(c echo.Context) error {
	return nil
}

func delete(c echo.Context) error {
	return nil
}

func routes() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Routes
	e.POST("/createRoute", create)
	e.GET("/findRoute", find)
	e.PUT("/updateRoute", update)
	e.DELETE("/deleteRoute", delete)
}

// Source: https://echo.labstack.com/guide/routing
