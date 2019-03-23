package packagename

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)


BasicAuthConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Validator is a function to validate BasicAuth credentials.
  // Required.
  Validator BasicAuthValidator

  // Realm is a string to define realm attribute of BasicAuth.
  // Default value "Restricted".
  Realm string
}


e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	if username == "joe" && password == "secret" {
		return true, nil
	}
	return false, nil
}))


e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))


DefaultBasicAuthConfig = BasicAuthConfig{
	Skipper: DefaultSkipper,
}
