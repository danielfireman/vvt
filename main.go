package main

import (
	"flag"
	"fmt"

	"github.com/danielfireman/vvt/Godeps/_workspace/src/github.com/labstack/echo"
	"github.com/danielfireman/vvt/todo"

	mw "github.com/danielfireman/vvt/Godeps/_workspace/src/github.com/labstack/echo/middleware"
)

var port = flag.Int("port", 8080, "Port to listen.")

func main() {
	flag.Parse()
	// Configuring server.
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Creating new todo store.
	s := todo.InMemoryStore()

	// Configuring API handlers.
	e.Post("/todo", todo.AddHandler(s))
	e.Get("/todo", todo.GetHandler(s))

	// Configuring HTTP handlers.
	e.Favicon("public/favicon.ico")
	e.Index("public/index.html")

	fmt.Printf("Server listening at localhost:%d\n", *port)
	e.Run(fmt.Sprintf(":%d", *port))
}
