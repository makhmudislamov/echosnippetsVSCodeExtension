package packagename

import (
	"net/http"	
	"github.com/labstack/echo"
)

var e = echo.New()

// Handlers
func create(c echo.Context) error {
}

func find(c echo.Context) error {
}

func update(c echo.Context) error {
}

func delete(c echo.Context) error {
}

// Routes
e.POST("/createRoute", create)
e.GET("/findRoute", find)
e.PUT("/updateRoute", update)
e.DELETE("/deleteRoute", delete)
