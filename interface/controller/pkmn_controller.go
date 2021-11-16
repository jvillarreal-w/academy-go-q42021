package controller

import (
	"net/http"

	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/interactor"
	u "github.com/jvillarreal-w/academy-go-q42021/utils"
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
		u.ErrorLogger.Printf("All Pokémon could not be fetched: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (pc *pokemonController) GetPokemonById(c Context) error {
	var p *model.Pokemon
	id := c.Param("id")

	p, err := pc.pokemonInteractor.GetById(p, id)

	if err != nil {
		u.ErrorLogger.Printf("Pokémon by ID could not be fetched: %v", err)
		return err
	}

	if p == nil {
		u.ErrorLogger.Printf("Pokémon by ID could not be found: %v", err)
		return c.JSON(http.StatusNotFound, p)
	}

	return c.JSON(http.StatusOK, p)
}
