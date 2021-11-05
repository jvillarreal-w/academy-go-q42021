package presenter

import (
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/presenter"
)

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemon(ps []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() presenter.PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemon(ps []*model.Pokemon) []*model.Pokemon {
	return ps
}
