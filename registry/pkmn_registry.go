package registry

import (
	"github.com/jvillarreal-w/academy-go-q42021/interface/controller"
	ir "github.com/jvillarreal-w/academy-go-q42021/interface/repository"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/interactor"
	ur "github.com/jvillarreal-w/academy-go-q42021/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository()
}
