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
}

func NewPokemonController(ps interactor.PokemonInteractor) PokemonController {
	return &pokemonController{ps}
}

func (pc *pokemonController) GetPokemon(c Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
