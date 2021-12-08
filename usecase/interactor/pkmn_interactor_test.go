package interactor

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jvillarreal-w/academy-go-q42021/common"
	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/jvillarreal-w/academy-go-q42021/usecase/repository"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
}

type mockErrorRepository struct {
}

func (m *mockRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	return []*model.Pokemon{
		{},
	}, nil
}

func (m *mockRepository) FindById(p []*model.Pokemon, id string) (*model.Pokemon, error) {
	return &model.Pokemon{}, nil
}

func (m *mockRepository) FindAllConcurrently(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {
	return []*model.Pokemon{
		{},
	}, nil
}

func (m *mockErrorRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	return nil, errors.New("")
}

func (m *mockErrorRepository) FindById(p []*model.Pokemon, id string) (*model.Pokemon, error) {
	return nil, errors.New("")
}

func (m *mockErrorRepository) FindAllConcurrently(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {
	return nil, errors.New("")
}

type (
	testCase struct {
		message       string
		repo          repository.PokemonRepository
		expectedError error
	}

	testCaseGet struct {
		expectedResult []*model.Pokemon
		testCase
	}

	testCaseFindById struct {
		expectedResult *model.Pokemon
		testCase
	}
)

var (
	getTestCases = []testCaseGet{
		{
			testCase: testCase{
				message: "Should return a slice of Pokemon",
				repo:    &mockRepository{},
			},
			expectedResult: []*model.Pokemon{
				{},
			},
		},
		{
			testCase: testCase{
				message:       "Should return an error",
				repo:          &mockErrorRepository{},
				expectedError: errors.New(""),
			},
		},
	}
	findByTestCase = []testCaseFindById{
		{
			expectedResult: &model.Pokemon{},
			testCase: testCase{
				message: "Should return a single Pokemon",
				repo:    &mockRepository{},
			},
		},
		{
			testCase: testCase{
				message:       "Should return an error",
				repo:          &mockErrorRepository{},
				expectedError: errors.New(""),
			},
		},
	}
)

func TestGet(t *testing.T) {
	for _, test := range getTestCases {
		mockInteractor := NewPokemonInteractor(test.repo)
		var p []*model.Pokemon
		p, err := mockInteractor.Get(p)

		if p != nil {
			assert.Equal(t, len(p), len(test.expectedResult), fmt.Sprintf("Expected %v elements in the response but got %v instead", len(test.expectedResult), len(p)))
		}

		if err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Errorf("Something unexpected occurred: %s", err)
			}
		}
	}
}

func TestGetById(t *testing.T) {
	for _, test := range findByTestCase {
		mockInteractor := NewPokemonInteractor(test.repo)
		var p []*model.Pokemon
		pkmnId := "1"
		pget, err := mockInteractor.GetById(p, pkmnId)

		if pget != nil {
			if pget.ID != test.expectedResult.ID {
				assert.Equal(t, test.expectedResult.ID, pget.ID, fmt.Sprintf("Should be id: %v but got id: %v instead", test.expectedResult.ID, pget.ID))
			}
		}

		if err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Errorf("Something unexpected occurred: %s", err)
			}
		}
	}
}

func TestGetConcurrently(t *testing.T) {
	for _, test := range getTestCases {
		mockInteractor := NewPokemonInteractor(test.repo)
		var p []*model.Pokemon
		p, err := mockInteractor.GetConcurrently(p, common.Even, 50, 5)

		if p != nil {
			assert.Equal(t, len(p), len(test.expectedResult), fmt.Sprintf("Expected %v elements in the response but got %v instead", len(test.expectedResult), len(p)))
		}

		if err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Errorf("Something unexpected occurred: %s", err)
			}
		}
	}
}
