package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/jvillarreal-w/academy-go-q42021/common"
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/interface/external"
	"github.com/jvillarreal-w/academy-go-q42021/interface/icontext"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/interactor"
	u "github.com/jvillarreal-w/academy-go-q42021/utils"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
	pokemonExternal   external.PokemonExternal
}

type PokemonController interface {
	GetPokemon(c icontext.IContext) error
	GetPokemonById(c icontext.IContext) error
	GetPokemonConcurrently(c icontext.IContext) error
}

func NewPokemonController(pi interactor.PokemonInteractor, pe external.PokemonExternal) PokemonController {
	return &pokemonController{pi, pe}
}

func (pc *pokemonController) GetPokemon(c icontext.IContext) error {
	var p []*model.Pokemon

	external_module := external.NewPokemonExternal()

	external_pkmn, err := external_module.GetExternalPokemon(p, c)
	if err != nil {
		u.ErrorLogger.Printf("External Pokemon could not be fetched: %s", err)
		return err
	}
	external_module.SaveExternalPokemon(external_pkmn)

	p, err = pc.pokemonInteractor.Get(p)
	if err != nil {
		u.ErrorLogger.Printf("All Pokémon could not be fetched: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (pc *pokemonController) GetPokemonById(c icontext.IContext) error {
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

func (pc *pokemonController) GetPokemonConcurrently(c icontext.IContext) error {
	var p []*model.Pokemon

	items, err := strconv.ParseInt(c.QueryParam(common.ItemsParam), 10, 64)
	if err != nil || items < 1 {
		return c.JSON(http.StatusBadRequest, u.ResponseBuilder(http.StatusBadRequest, "query parameter 'items' must be numeric and greater than 0"))
	}

	itemsWorker, err := strconv.ParseInt(c.QueryParam(common.ItemsPerWorkerParam), 10, 64)
	if err != nil || itemsWorker < 1 {
		return c.JSON(http.StatusBadRequest, u.ResponseBuilder(http.StatusBadRequest, "query parameter 'items_per_worker' must be numeric and greater than 0"))
	}

	if itemsWorker > items {
		return c.JSON(http.StatusBadRequest, u.ResponseBuilder(http.StatusBadRequest, "'items_per_worker' parameter shouldn't have a greater value than 'items'"))
	}

	t := strings.ToLower(c.QueryParam(common.TypeParam))
	if t != "" && t != common.Odd && t != common.Even {
		return c.JSON(http.StatusBadRequest, u.ResponseBuilder(http.StatusBadRequest, "query parameter 'type' only supports 'even' and 'odd'"))
	}

	p, err = pc.pokemonInteractor.GetConcurrently(p, t, items, itemsWorker)
	if err != nil {
		u.ErrorLogger.Printf("Pokemon could not be fetched concurrently: %s", err)
		return c.JSON(http.StatusInternalServerError, u.ResponseBuilder(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, p)
}
