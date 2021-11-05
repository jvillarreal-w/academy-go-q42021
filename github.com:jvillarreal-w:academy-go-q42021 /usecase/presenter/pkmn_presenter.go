package presenter

import "github.com/jvillarreal-w/academy-go-q42021/domain/model"

type PokemonPresenter interface {
	ResponsePokemon(u []*model.Pokemon) []*model.Pokemon
}
