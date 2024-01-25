package context

import "github.com/labstack/echo/v4"

type Custom struct {
	echo.Context
}

func (c *Custom) CurrentURL() string {
	return c.Request().URL.Path
}
