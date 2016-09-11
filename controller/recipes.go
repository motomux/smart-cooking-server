package controller

import (
	"net/http"
	"strconv"

	tarantool "github.com/tarantool/go-tarantool"

	"github.com/julienschmidt/httprouter"
	"github.com/motomux/smart-cooking-server/service"
)

// RecipesCtrl is a controller for user service
type RecipesCtrl struct {
	Svc service.RecipesSvcInterface
}

// NewRecipesCtrl initiates RecipesCtrl
func NewRecipesCtrl(client *tarantool.Connection) *RecipesCtrl {
	return &RecipesCtrl{
		Svc: service.NewRecipesSvc(client),
	}
}

// GetOne parses http request, calls service and writes http response
func (u *RecipesCtrl) GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	recipeID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		respondErr(w, r, http.StatusBadRequest, err)
		return
	}

	Recipe, err := u.Svc.GetOne(recipeID)
	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, err)
		return
	}

	respond(w, r, http.StatusOK, Recipe)
}
