package packageName

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
		fmt.Println("rendering Hello World")
	})
	e.Logger.Fatal(e.Start(":1323"))
}