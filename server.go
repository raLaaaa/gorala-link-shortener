package main

import (
	"io"
	"os"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	a := controllers.AdminController{}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	g := e.Group("/admin")
	g.POST("/delete/:linkId", l.DeleteLink)
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == os.Getenv("link-admin-name") && password == os.Getenv("link-admin-pw") {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/main", a.ShowAdmin)

	e.POST("/shorten", l.ShortenLink)
	e.GET("/", l.RedirectShortendLink)
	e.GET("/:shortLink", l.RedirectShortendLink)
	e.Logger.Fatal(e.Start(":1333"))
}
