package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/danielfireman/vvt/todo"
	"github.com/labstack/echo"

	mw "github.com/labstack/echo/middleware"
)

var port = flag.Int("port", 8999, "Port to listen.")

type item struct {
	Desc string
}

func main() {
	flag.Parse()
	// Configuring server.
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Creating new todo store.
	s := todo.NewStore()

	// Handlers.
	e.Post("/todo", func(c *echo.Context) error {
		var t item
		if err := c.Bind(&t); err != nil {
			return c.JSON(http.StatusPreconditionFailed, err.Error())
		}
		s.Add(t.Desc)
		return c.JSON(http.StatusCreated, t)
	})

	e.Get("/todo", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.List())
	})
	fmt.Printf("Server listening at localhost:%d\n", *port)
	e.Run(fmt.Sprintf(":%d", *port))
}
