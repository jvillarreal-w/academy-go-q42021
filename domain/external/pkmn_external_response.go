package external

import "github.com/jvillarreal-w/academy-go-q42021/domain/model"

type PokemonExternalResponse struct {
	Count   uint64          `json:"count"`
	Results []model.Pokemon `json:"results"`
}
