package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/488Ques/aws-demo/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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

	// Add default data passed to templates
	td, ok := data.(*templateData)
	if !ok {
		td = new(templateData)
	}
	err := addDefaultData(td, c)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "base.html", td)
}

func main() {
	e := echo.New()

	// Server static files
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
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	secret := "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge"
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))

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
