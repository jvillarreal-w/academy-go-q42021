package interactor

import (
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
}

type PokemonInteractor interface {
	Get(p []*model.Pokemon) ([]*model.Pokemon, error)
	GetById(p []*model.Pokemon, id string) (*model.Pokemon, error)
	GetConcurrently(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository) PokemonInteractor {
	return &pokemonInteractor{r}
}

func (ps *pokemonInteractor) Get(p []*model.Pokemon) ([]*model.Pokemon, error) {
	return ps.PokemonRepository.FindAll(p)
}

func (pi *pokemonInteractor) GetById(p []*model.Pokemon, id string) (*model.Pokemon, error) {
	return pi.PokemonRepository.FindById(p, id)
}

func (pi *pokemonInteractor) GetConcurrently(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {

	p, err := pi.PokemonRepository.FindAllConcurrently(p, t, items, itemsWorker)
	if err != nil {
		return nil, err
	}
	return p, nil
}
