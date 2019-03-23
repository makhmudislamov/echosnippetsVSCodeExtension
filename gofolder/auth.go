package packagename

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"net/http"
)

func ValidateUser(username, password string, c echo.Context) (bool, error) {
	if username == "joe" && password == "secret" {
		return true, nil
	}
	return false, nil
}
