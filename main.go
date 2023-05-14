package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/488Ques/aws-demo/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found: " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	e := echo.New()

	e.Static("/static", "static")

	// Renderer
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("views/home.html", "views/base.html"))
	templates["login.html"] = template.Must(template.ParseFiles("views/login.html", "views/base.html"))
	templates["inventory.html"] = template.Must(template.ParseFiles("views/inventory.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Middlewares
	e.Use(middleware.Recover())

	// Routes
	Routes(e)

	// Init database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	err = dbGorm.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database!")

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
