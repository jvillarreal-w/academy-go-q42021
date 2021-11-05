package repository

import "github.com/jvillarreal-w/academy-go-q42021/domain/model"

type PokemonRepository interface {
	//FindById(p []*model.Pokemon) (*model.Pokemon, error)
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}
