package external

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
	SaveExternalPokemon([]*model.Pokemon)
}

func NewPokemonExternal() PokemonExternal {
	return &pokemonExternal{}
}

func (pe *pokemonExternal) GetExternalPokemon(p []*model.Pokemon, c context.Context) ([]*model.Pokemon, error) {
	request, err := http.Get(fmt.Sprintf("%v?limit=50", url))

	if err != nil {
		u.ErrorLogger.Println("External resource is not reachable")
		return nil, err
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		u.ErrorLogger.Println("Response body could not be read")
		return nil, err
	}

	response := external.PokemonExternalResponse{}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	for i, val := range response.Results {
		id := uint64(i + 1)

		details_request, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", id))

		if err != nil {
			u.ErrorLogger.Printf("External details are not reachable: %s", err)
			return nil, err
		}

		details_body, err := ioutil.ReadAll(details_request.Body)

		if err != nil {
			u.ErrorLogger.Println("Details body could not be read")
			return nil, err
		}

		details_response := external.PokemonExternalDetails{}

		if err := json.Unmarshal(details_body, &details_response); err != nil {
			u.ErrorLogger.Printf("Details could not be unmarshalled: %v", err)
			return nil, err
		}

		if len(details_response.Types) < 2 {
			pkmn := &model.Pokemon{
				ID:          id,
				Name:        val.Name,
				Generation:  getGeneration(id),
				PrimaryType: details_response.Types[0].TypeStructure.TypeName,
				Stats: model.Stats{
					HP:             details_response.Stats[0].BaseStat,
					Attack:         details_response.Stats[1].BaseStat,
					Defense:        details_response.Stats[2].BaseStat,
					SpecialAttack:  details_response.Stats[3].BaseStat,
					SpecialDefense: details_response.Stats[4].BaseStat,
					Speed:          details_response.Stats[5].BaseStat,
					BaseStatTotal:  getBST(details_response.Stats),
				},
			}
			u.GeneralLogger.Println(pkmn)
			p = append(p, pkmn)
		} else {
			pkmn := &model.Pokemon{
				ID:            id,
				Name:          val.Name,
				Generation:    getGeneration(id),
				PrimaryType:   details_response.Types[0].TypeStructure.TypeName,
				SecondaryType: details_response.Types[1].TypeStructure.TypeName,
				Stats: model.Stats{
					HP:             details_response.Stats[0].BaseStat,
					Attack:         details_response.Stats[1].BaseStat,
					Defense:        details_response.Stats[2].BaseStat,
					SpecialAttack:  details_response.Stats[3].BaseStat,
					SpecialDefense: details_response.Stats[4].BaseStat,
					Speed:          details_response.Stats[5].BaseStat,
					BaseStatTotal:  getBST(details_response.Stats),
				},
			}
			u.GeneralLogger.Println(pkmn)
			p = append(p, pkmn)
		}
	}
	return p, nil
}

func (pe *pokemonExternal) SaveExternalPokemon(p []*model.Pokemon) {
	csvFile, err := os.Create("pkmn.csv")

	if err != nil {
		u.ErrorLogger.Printf("Failed creating CSV file: %s", err)
	}

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Write([]string{"ID", "Name", "Primary Type", "Secondary Type", "Generation", "HP", "Attack", "Defense", "Special Attack", "Special Defense", "Speed", "Base Stat Total"})
	for _, pkmnRow := range p {
		_ = csvWriter.Write(pkmnRow.ToStringSlice())
	}
	csvWriter.Flush()
	csvFile.Close()
}

func getBST(s []external.PokemonExternalStats) uint64 {
	bst := 0
	for _, v := range s {
		bst += int(v.BaseStat)
	}
	return uint64(bst)
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
