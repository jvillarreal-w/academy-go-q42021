package router

import (
	"github.com/jvillarreal-w/academy-go-q42021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemon", func(context echo.Context) error { return c.Pokemon.GetPokemon(context) })
	e.GET("/pokemon/:id", func(context echo.Context) error { return c.Pokemon.GetPokemonById(context) })
	e.GET("/pokemon/concurrent", func(context echo.Context) error { return c.Pokemon.GetPokemonConcurrently(context) })

	return e
}
