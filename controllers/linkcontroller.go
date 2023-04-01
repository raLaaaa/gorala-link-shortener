package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raLaaaa/gorala-link-shortener/models.go"
	"github.com/raLaaaa/gorala-link-shortener/services"
)

type LinkController struct{}

type LinkDTO struct {
	Link string `json:"link"`
}

func (t *LinkController) RedirectShortendLink(c echo.Context) error {
	shortLink := c.Param("shortLink")

	if shortLink == "" {
		return c.Render(http.StatusOK, "main", "")
	}

	dbService := services.DatabaseService{}
	link, err := dbService.FindLinkByShortLink(shortLink)

	if err != nil || link == nil {
		return c.Redirect(302, "/")
	}

	dbService.IncreaseNumbersOfClicked(shortLink)
	return c.Redirect(302, link.Link)
}

func (t *LinkController) ShortenLink(c echo.Context) error {

	linkDTO := new(LinkDTO)

	if err := c.Bind(linkDTO); err != nil {
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}

	link := models.Link{
		Link: linkDTO.Link,
	}

	host := c.Request().Host

	dbService := services.DatabaseService{}
	dbService.CreateShortLink(&link, host)

	return c.JSON(http.StatusCreated, link)
}
