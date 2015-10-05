package main

import (
	"flag"
	"fmt"

	"github.com/danielfireman/vvt/todo"
	"github.com/labstack/echo"

	mw "github.com/labstack/echo/middleware"
)

var port = flag.Int("port", 8999, "Port to listen.")

func main() {
	flag.Parse()
	// Configuring server.
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Creating new todo store.
	s := todo.NewStore()

	// Configuring handlers.
	e.Post("/todo", todo.NewAddHandler(s))
	e.Get("/todo", todo.NewGetHandler(s))

	fmt.Printf("Server listening at localhost:%d\n", *port)
	e.Run(fmt.Sprintf(":%d", *port))
}
