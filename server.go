package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/raLaaaa/gorala-link-shortener/controllers"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (tem *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return tem.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	l := controllers.LinkController{}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "static")

	e.POST("/shorten", l.ShortenLink)
	e.GET("/", l.RedirectShortendLink)
	e.GET("/:shortLink", l.RedirectShortendLink)
	e.Logger.Fatal(e.Start(":1333"))
}
