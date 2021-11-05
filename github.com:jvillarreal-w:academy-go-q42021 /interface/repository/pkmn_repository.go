package repository

import (
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/repository"
)

type pokemonRepository struct {
}

func NewPokemonRepository() repository.PokemonRepository {
	return &pokemonRepository{}
}

func (pr *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	return p, nil
}

// func (pr *pokemonRepository) FindById(p *model.Pokemon) (*model.Pokemon, error) {
// 	return p, nil
// }
