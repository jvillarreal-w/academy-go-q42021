package repository

import "github.com/jvillarreal-w/academy-go-q42021/domain/model"

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
	FindById(p []*model.Pokemon, id string) (*model.Pokemon, error)
	FindAllConcurrently(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error)
}
