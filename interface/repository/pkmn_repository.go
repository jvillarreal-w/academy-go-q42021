package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/jvillarreal-w/academy-go-q42021/common"
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/repository"
	u "github.com/jvillarreal-w/academy-go-q42021/utils"
)

type pokemonRepository struct {
	FilePath string
}

var lock = new(sync.Mutex)

func NewPokemonRepository(path string) repository.PokemonRepository {
	return &pokemonRepository{FilePath: path}
}

func openInternalDataSource(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)

	if err != nil {
		u.ErrorLogger.Printf("Data source could not be opened: %v", err)
		return nil, err
	}
	fmt.Println("Successfully opened CSV file")

	return f, err
}

func getInternalDataSourceReader(file *os.File) (*csv.Reader, error) {
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	if _, err := r.Read(); err != nil {
		u.ErrorLogger.Print("Line could not be read.")
		return nil, err
	}
	return r, nil
}

func readData(reader *csv.Reader) ([][]string, error) {
	rows, err := reader.ReadAll()

	return rows, err
}

func (pr *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	file, err := openInternalDataSource(pr.FilePath)

	if err != nil {
		u.ErrorLogger.Println("File could not be open.")
		return nil, err
	}

	reader, err := getInternalDataSourceReader(file)

	if err != nil {
		u.ErrorLogger.Println("Reader could not be obtained.")
		return nil, err
	}

	rows, err := readData(reader)

	if err != nil {
		u.ErrorLogger.Println("Records could not be read.")
		return nil, err
	}

	return parsePokemon(p, rows)
}

func parsePokemon(p []*model.Pokemon, rows [][]string) ([]*model.Pokemon, error) {
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

func worker(r *csv.Reader, t string, itemsWorker int64, results chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()
	var lines int64
	for {
		if lines == itemsWorker {
			break
		}
		lock.Lock()
		line, err := r.Read()
		lock.Unlock()
		if err == io.EOF {
			break
		}
		if len(line) != 12 {
			continue
		}

		pid, err := strconv.ParseUint(line[0], 10, 32)
		if err != nil {
			u.ErrorLogger.Println("Failed to convert ID to uint.")
		}
		if oddEvenCriteriaMet(t, pid) {
			results <- line
			lines++
		}
	}
}

func oddEvenCriteriaMet(t string, pid uint64) bool {
	switch t {
	case common.Odd:
		return int(pid)%2 != 0
	case common.Even:
		return int(pid)%2 == 0
	default:
		u.ErrorLogger.Println("Default case in odd/even criteria was entered.")
		return true
	}

}

func readDataConcurrently(fileName string, p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {
	var result [][]string
	wg := new(sync.WaitGroup)
	lines := make(chan []string, items)
	workers := items / itemsWorker

	f, err := openInternalDataSource(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r, err := getInternalDataSourceReader(f)

	if err != nil {
		return nil, err
	}

	for w := int64(0); w < workers; w++ {
		wg.Add(1)
		go worker(r, t, itemsWorker, lines, wg)
	}

	go func(lines chan []string) {
		wg.Wait()
		close(lines)
	}(lines)

	for line := range lines {
		result = append(result, line)
	}

	return parsePokemon(p, result)
}

func (pr *pokemonRepository) FindAllConcurrently(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {
	return readDataConcurrently(pr.FilePath, p, t, items, itemsWorker)
}
