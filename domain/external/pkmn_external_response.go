package external

import (
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
)

type PokemonExternalResponse struct {
	Count   uint64          `json:"count"`
	Results []model.Pokemon `json:"results"`
}

type PokemonExternalStats struct {
	BaseStat uint64 `json:"base_stat"`
}

type PokemonExternalTypes struct {
	TypeStructure PokemonExternalTypeProperty `json:"type"`
}

type PokemonExternalTypeProperty struct {
	TypeName string `json:"name"`
}

type PokemonExternalDetails struct {
	Stats []PokemonExternalStats `json:"stats"`
	Types []PokemonExternalTypes `json:"types"`
}
