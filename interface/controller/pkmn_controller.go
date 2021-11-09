package controller

import (
	"net/http"

	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemon(c Context) error
	GetPokemonById(c Context) error
}

func NewPokemonController(pi interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetPokemon(c Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (pc *pokemonController) GetPokemonById(c Context) error {
	var p *model.Pokemon
	id := c.Param("id")

	p, err := pc.pokemonInteractor.GetById(p, id)

	if err != nil {
		return err
	}

	if p == nil {
		return c.JSON(http.StatusNotFound, p)
	}

	return c.JSON(http.StatusOK, p)
}
