package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/repository"
)

type pokemonRepository struct {
}

func NewPokemonRepository() repository.PokemonRepository {
	return &pokemonRepository{}
}

func readInternalDataSource(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully opened CSV file")
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return rows[1:], nil
}

func (pr *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	rows, err := readInternalDataSource("pkmn.csv")

	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		id, err := strconv.ParseUint(row[0], 10, 32)

		if err != nil {
			return nil, err
		}

		pokemon := &model.Pokemon{
			ID:   id,
			Name: row[1],
		}

		p = append(p, pokemon)
	}

	return p, nil
}

func (pr *pokemonRepository) FindById(p *model.Pokemon, id string) (*model.Pokemon, error) {
	rows, err := readInternalDataSource("pkmn.csv")

	if err != nil {
		return nil, err
	}

	pkmnId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		dataSourceId, err := strconv.ParseUint(row[0], 10, 32)

		if err != nil {
			return nil, err
		}

		if dataSourceId != pkmnId {
			continue
		}

		p = &model.Pokemon{
			ID:   dataSourceId,
			Name: row[1],
		}

		break
	}

	return p, nil
}
