package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/repository"
	u "github.com/jvillarreal-w/academy-go-q42021/utils"
)

type pokemonRepository struct {
	FilePath string
}

func NewPokemonRepository(path string) repository.PokemonRepository {
	return &pokemonRepository{FilePath: path}
}

func readInternalDataSource(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		u.ErrorLogger.Printf("Data source could not be opened: %v", err)
		return nil, err
	}
	fmt.Println("Successfully opened CSV file")
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		u.ErrorLogger.Printf("Data source content could not be read: %v", err)
		return nil, err
	}

	return rows[1:], nil
}

func (pr *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	rows, err := readInternalDataSource(pr.FilePath)

	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		id, err := strconv.ParseUint(row[0], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute ID: %v", err)
			return nil, err
		}

		gen, err := strconv.ParseUint(row[4], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute Generation: %v", err)
			return nil, err
		}

		hp, err := strconv.ParseUint(row[5], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute HP: %v", err)
			return nil, err
		}

		atk, err := strconv.ParseUint(row[6], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute Attack: %v", err)
			return nil, err
		}

		def, err := strconv.ParseUint(row[7], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute Defense: %v", err)
			return nil, err
		}

		spa, err := strconv.ParseUint(row[8], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute Special Attack: %v", err)
			return nil, err
		}

		spd, err := strconv.ParseUint(row[9], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute Special Defense: %v", err)
			return nil, err
		}

		spe, err := strconv.ParseUint(row[10], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute Speed: %v", err)
			return nil, err
		}

		bst, err := strconv.ParseUint(row[11], 10, 32)

		if err != nil {
			u.ErrorLogger.Printf("Invalid attribute Base Stat Total: %v", err)
			return nil, err
		}

		pokemon := &model.Pokemon{
			ID:            id,
			Name:          row[1],
			PrimaryType:   row[2],
			SecondaryType: row[3],
			Generation:    gen,
			Stats: model.Stats{
				HP:             hp,
				Attack:         atk,
				Defense:        def,
				SpecialAttack:  spa,
				SpecialDefense: spd,
				Speed:          spe,
				BaseStatTotal:  bst,
			},
		}

		p = append(p, pokemon)
	}

	return p, nil
}

func (pr *pokemonRepository) FindById(p []*model.Pokemon, id string) (*model.Pokemon, error) {
	pkmnId, _ := strconv.ParseUint(id, 10, 32)

	pkmnList, err := pr.FindAll(p)

	if err != nil {
		u.ErrorLogger.Println("Pokemon list could not be fetched.")
		return nil, err
	}

	for _, pkmn := range pkmnList {
		dataSourceId := pkmn.ID

		if dataSourceId != pkmnId {
			continue
		}
		return pkmn, nil
	}
	return nil, err
}
