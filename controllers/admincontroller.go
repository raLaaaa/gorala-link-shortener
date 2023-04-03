package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raLaaaa/gorala-link-shortener/services"
)

type AdminController struct{}

func (a *AdminController) ShowAdmin(c echo.Context) error {
	dbService := services.DatabaseService{}
	links, err := dbService.FindAllLinks()

	if err != nil {
		return c.Render(http.StatusOK, "admin", err)
	}

	return c.Render(http.StatusOK, "admin", links)
}
