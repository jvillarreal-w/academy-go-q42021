package registry

import (
	"github.com/jvillarreal-w/academy-go-q42021/common"
	"github.com/jvillarreal-w/academy-go-q42021/interface/controller"
	"github.com/jvillarreal-w/academy-go-q42021/interface/external"
	ir "github.com/jvillarreal-w/academy-go-q42021/interface/repository"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/interactor"
	ur "github.com/jvillarreal-w/academy-go-q42021/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor(), r.NewPokemonExternal())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository(common.InternalDataSourcePath)
}

func (r *registry) NewPokemonExternal() external.PokemonExternal {
	return external.NewPokemonExternal()
}
