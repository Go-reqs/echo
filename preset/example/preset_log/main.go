package main

import (
	"github.com/go-reqs/echo/preset"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()
	preset.Init(e)

	// Route => handler
	e.GET("/*", func(c echo.Context) error {
		c.Logger().Infof("recieve request: %s", c.Request().RequestURI)
		return c.JSON(http.StatusOK, "Hi!")
	})

	e.Logger.Fatal(e.Start(":4000"))
}
