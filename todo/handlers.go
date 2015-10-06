package todo

import (
	"net/http"

	"github.com/danielfireman/vvt/Godeps/_workspace/src/github.com/labstack/echo"
)

// Handler for the add todo operation.
func AddHandler(s *store) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var t item
		if err := c.Bind(&t); err != nil {
			return c.JSON(http.StatusPreconditionFailed, err.Error())
		}
		s.Add(&t)
		return c.JSON(http.StatusCreated, t)
	}
}

// Handler for the get todo operation.
func GetHandler(s *store) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.List())
	}
}
