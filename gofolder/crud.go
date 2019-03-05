package packagename

import (
        "github.com/labstack/echo"
        "net/http"
)

// Handlers
func create(c echo.Context) error {

}

func find(c echo.Context) error {
}

func update(c echo.Context) error {
}

func delete(c echo.Context) error {
}

func main() {

        e := echo.New()
        // Routes
        e.POST("/createRoute", create)
        e.GET("/findRoute", find)
        e.PUT("/updateRoute", update)
        e.DELETE("/deleteRoute", delete)
}