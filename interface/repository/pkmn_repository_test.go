package repository

import (
	"fmt"
	"testing"

	"github.com/jvillarreal-w/academy-go-q42021/common"
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestPokemonRepository_FindAll(t *testing.T) {
	pr := NewPokemonRepository(common.TestDataSourcePath)
	rows, _ := readInternalDataSource(common.TestDataSourcePath)
	var p []*model.Pokemon

	p, _ = pr.FindAll(p)

	assert.Equal(t, len(p), len(rows), fmt.Sprintf("Expected %v but got %v", len(rows), len(p)))
}

func TestPokemonRepository_FindById(t *testing.T) {
	expected := "bulbasaur"
	pr := NewPokemonRepository(common.TestDataSourcePath)
	var p []*model.Pokemon
	pkmn, _ := pr.FindAll(p)

	pid, _ := pr.FindById(pkmn, "1")

	assert.Equal(t, pid.Name, expected, fmt.Sprintf("Expected %v but got %v", pid.Name, expected))
}
