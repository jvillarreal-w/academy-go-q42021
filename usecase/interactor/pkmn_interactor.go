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
	GetById(p *model.Pokemon, id string) (*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository) PokemonInteractor {
	return &pokemonInteractor{r}
}

func (ps *pokemonInteractor) Get(p []*model.Pokemon) ([]*model.Pokemon, error) {
	p, err := ps.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pi *pokemonInteractor) GetById(p *model.Pokemon, id string) (*model.Pokemon, error) {
	p, err := pi.PokemonRepository.FindById(p, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
