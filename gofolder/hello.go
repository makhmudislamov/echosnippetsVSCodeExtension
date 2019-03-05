package packagename

import (
        "fmt"
        "github.com/labstack/echo"
        "net/http"
)

func main() {
        e := echo.New()

        e.GET("/", func(c echo.Context) error {
                return c.String(http.StatusOK, "Hello, World!")
                // fmt.Println("rendering Hello World")
        })

        e.Logger.Fatal(e.Start(":3000"))
}