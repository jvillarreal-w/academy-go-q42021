package controller

import (
	"net/http"
	"strconv"

	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/interface/context"
	"github.com/jvillarreal-w/academy-go-q42021/interface/external"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/interactor"
	u "github.com/jvillarreal-w/academy-go-q42021/utils"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
	pokemonExternal   external.PokemonExternal
}

type PokemonController interface {
	GetPokemon(c context.Context) error
	GetPokemonById(c context.Context) error
}

func NewPokemonController(pi interactor.PokemonInteractor, pe external.PokemonExternal) PokemonController {
	return &pokemonController{pi, pe}
}

func (pc *pokemonController) GetPokemon(c context.Context) error {
	var p []*model.Pokemon

	external_pkmn, _ := external.NewPokemonExternal().GetExternalPokemon(p, c)
	external.NewPokemonExternal().SaveExternalPokemon(external_pkmn)

	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
		u.ErrorLogger.Printf("All Pokémon could not be fetched: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (pc *pokemonController) GetPokemonById(c context.Context) error {
	var p []*model.Pokemon
	id := c.Param("id")
	// Checking ID validity.
	_, err := strconv.Atoi(id)

	if err != nil {
		u.ErrorLogger.Printf("Invalid Pokemon ID: %v", err)
		return err
	}

	pkmn, err := pc.pokemonInteractor.GetById(p, id)

	if pkmn == nil {
		u.ErrorLogger.Printf("Pokémon could not be found by ID: %v", err)
		return c.JSON(http.StatusNotFound, p)
	}

	return c.JSON(http.StatusOK, pkmn)
}
