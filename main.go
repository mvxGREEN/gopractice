package main

import (
  "html/template"
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "io"
)

type IndexData struct {
	Host              string
	EscapedMessageVar string
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "layout", &IndexData{Host: c.Request().Host, EscapedMessageVar: "{{message}}"})
}

func main() {
  e := echo.New()

  e.Use(middleware.Recover())
	e.Use(middleware.Logger())


  t := &Template{
    templates: template.Must(template.ParseFiles("views/layout.html", "views/styles.html", "views/content.html")),
  }
  e.Renderer = t

	e.Static("/", "public")
	e.GET("/", Index)

  e.Logger.Fatal(e.Start(":80"))
}
