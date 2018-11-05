package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/hoge", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
