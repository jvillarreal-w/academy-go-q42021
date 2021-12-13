package external

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jvillarreal-w/academy-go-q42021/domain/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetExternalPokemon(t *testing.T) {
	expectedLength := 50
	var res []*model.Pokemon
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/pokemon", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	pe := NewPokemonExternal()

	res, _ = pe.GetExternalPokemon(res, c)

	assert.Equal(t, len(res), expectedLength, fmt.Sprintf("Expected %v elements in the response but got %v", expectedLength, len(res)))
}
