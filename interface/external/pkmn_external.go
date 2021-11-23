package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jvillarreal-w/academy-go-q42021/domain/external"
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/interface/context"
	u "github.com/jvillarreal-w/academy-go-q42021/utils"
)

const url = "https://pokeapi.co/api/v2/pokemon"

type pokemonExternal struct {
}

type PokemonExternal interface {
	GetExternalPokemon(p []*model.Pokemon, c context.Context) ([]*model.Pokemon, error)
	//SaveExternalPokemon([]*model.Pokemon)
}

func NewPokemonExternal() PokemonExternal {
	return &pokemonExternal{}
}

func (pe *pokemonExternal) GetExternalPokemon(p []*model.Pokemon, c context.Context) ([]*model.Pokemon, error) {
	request, err := http.Get(fmt.Sprintf("%v?limit=151", url))

	if err != nil {
		u.ErrorLogger.Println("External resource is not reachable")
		return nil, err
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		u.ErrorLogger.Println("Response could not be read")
		return nil, err
	}

	response := external.PokemonExternalResponse{}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	for i, val := range response.Results {
		id := uint64(i + 1)

		pkmn := &model.Pokemon{
			ID:         id,
			Name:       val.Name,
			Generation: getGeneration(id),
		}

		u.GeneralLogger.Printf("%v", pkmn.ID)

		p = append(p, pkmn)
	}
	return p, nil
}

func getGeneration(id uint64) uint64 {
	switch {
	case id <= 151:
		return 1
	case id <= 251:
		return 2
	case id <= 386:
		return 3
	case id <= 493:
		return 4
	case id <= 649:
		return 5
	case id <= 721:
		return 6
	case id <= 809:
		return 7
	case id <= 898:
		return 8
	default:
		return 0
	}
}
