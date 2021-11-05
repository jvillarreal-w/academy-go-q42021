package router

import (
	"github.com/jvillarreal-w/academy-go-q42021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemon", func(context echo.Context) error { return c.GetPokemon(context) })

	return e
}
