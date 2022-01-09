package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var myarr [3]string

		myarr[0] = "Hello"
		myarr[1] = "From"
		myarr[2] = "GoLang"

		return c.JSON(http.StatusOK, myarr)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
