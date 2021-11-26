package repository

import (
	"testing"

	"github.com/jvillarreal-w/academy-go-q42021/common"
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
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
	expected := "pikachu"
	pr := NewPokemonRepository(common.TestDataSourcePath)
	var p []*model.Pokemon
	pkmn, _ := pr.FindAll(p)

	if pid, _ := pr.FindById(pkmn, "25"); pid.Name != expected {
		t.Errorf("Expected: %s but got: %s", expected, pid.Name)
	}

}
