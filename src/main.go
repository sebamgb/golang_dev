package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// intancia de echo
	e := echo.New()

	// ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hola mundo")
	})

	// servidor
	e.Logger.Fatal(e.Start(":1323"))
}
