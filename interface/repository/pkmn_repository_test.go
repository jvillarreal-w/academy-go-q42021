package repository

import (
	"strings"
	"testing"

	"github.com/jvillarreal-w/academy-go-q42021/common"
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	u "github.com/jvillarreal-w/academy-go-q42021/utils"
)

func TestPokemonRepository_FindAll(t *testing.T) {
	pr := NewPokemonRepository(common.TestDataSourcePath)
	rows, _ := readInternalDataSource(common.TestDataSourcePath)
	var p []*model.Pokemon

	if p, _ = pr.FindAll(p); len(p) != len(rows) {
		t.Errorf("Expected: %v but got: %v", len(rows), len(p))
	}
}

func TestPokemonRepository_FindById(t *testing.T) {
	expected := "bulbasaur"
	pr := NewPokemonRepository(common.TestDataSourcePath)
	var p []*model.Pokemon
	pkmn, _ := pr.FindAll(p)

	pid, _ := pr.FindById(pkmn, "1")
	u.GeneralLogger.Printf("%+v", pid)
	if strings.Compare(pid.Name, expected) != 0 {
		t.Errorf("Expected: %s but got: %s", expected, pid.Name)
	}
}
