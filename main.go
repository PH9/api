package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ndidplatform/api/user"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/v1/users", user.GetUser)

	e.Logger.Fatal(e.Start(":2018"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "What's up!!! I'm NDID platform")
}
