package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/jvillarreal-w/academy-go-q42021/common"
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
	rows, err := readInternalDataSource(common.InternalDataSourcePath)

	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		id, err := strconv.ParseUint(row[0], 10, 32)

		if err != nil {
			return nil, err
		}

		gen, err := strconv.ParseUint(row[4], 10, 32)

		if err != nil {
			return nil, err
		}

		hp, err := strconv.ParseUint(row[5], 10, 32)

		if err != nil {
			return nil, err
		}

		atk, err := strconv.ParseUint(row[6], 10, 32)

		if err != nil {
			return nil, err
		}

		def, err := strconv.ParseUint(row[7], 10, 32)

		if err != nil {
			return nil, err
		}

		spa, err := strconv.ParseUint(row[8], 10, 32)

		if err != nil {
			return nil, err
		}

		spd, err := strconv.ParseUint(row[9], 10, 32)

		if err != nil {
			return nil, err
		}

		spe, err := strconv.ParseUint(row[10], 10, 32)

		if err != nil {
			return nil, err
		}

		bst, err := strconv.ParseUint(row[11], 10, 32)

		if err != nil {
			return nil, err
		}

		pokemon := &model.Pokemon{
			ID:             id,
			Name:           row[1],
			PrimaryType:    row[2],
			SecondaryType:  row[3],
			Generation:     gen,
			HP:             hp,
			Attack:         atk,
			Defense:        def,
			SpecialAttack:  spa,
			SpecialDefense: spd,
			Speed:          spe,
			BaseStatTotal:  bst,
		}

		p = append(p, pokemon)
	}

	return p, nil
}

func (pr *pokemonRepository) FindById(p *model.Pokemon, id string) (*model.Pokemon, error) {
	rows, err := readInternalDataSource(common.InternalDataSourcePath)

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

		gen, err := strconv.ParseUint(row[4], 10, 32)

		if err != nil {
			return nil, err
		}

		hp, err := strconv.ParseUint(row[5], 10, 32)

		if err != nil {
			return nil, err
		}

		atk, err := strconv.ParseUint(row[6], 10, 32)

		if err != nil {
			return nil, err
		}

		def, err := strconv.ParseUint(row[7], 10, 32)

		if err != nil {
			return nil, err
		}

		spa, err := strconv.ParseUint(row[8], 10, 32)

		if err != nil {
			return nil, err
		}

		spd, err := strconv.ParseUint(row[9], 10, 32)

		if err != nil {
			return nil, err
		}

		spe, err := strconv.ParseUint(row[10], 10, 32)

		if err != nil {
			return nil, err
		}

		bst, err := strconv.ParseUint(row[11], 10, 32)

		if err != nil {
			return nil, err
		}

		p = &model.Pokemon{
			ID:             dataSourceId,
			Name:           row[1],
			PrimaryType:    row[2],
			SecondaryType:  row[3],
			Generation:     gen,
			HP:             hp,
			Attack:         atk,
			Defense:        def,
			SpecialAttack:  spa,
			SpecialDefense: spd,
			Speed:          spe,
			BaseStatTotal:  bst,
		}

		break
	}

	return p, nil
}
