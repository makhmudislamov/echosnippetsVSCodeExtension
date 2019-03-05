package packagename

import (
        "github.com/labstack/echo"
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

func route() {

        e := echo.New()
        // Routes
        e.POST("/createRoute", create)
        e.GET("/findRoute", find)
        e.PUT("/updateRoute", update)
        e.DELETE("/deleteRoute", delete)
}