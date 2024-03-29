package main

import (
	"profileGenerator/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/:seed", handlers.GetProfile)
	e.Logger.Fatal(e.Start(":6942"))

}
