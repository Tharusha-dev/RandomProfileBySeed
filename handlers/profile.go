package handlers

import (
	// "net/http"
	"net/http"
	"profileGenerator/utils"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {

	seed := c.Param("seed")
	email, address, username, _ := utils.GetProfileFromSeed(seed)

	return c.HTML(http.StatusOK, "<strong>Hello, World!</strong> "+email+address+username)

}
