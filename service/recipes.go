package service

import (
	"github.com/motomux/smart-cooking-server/resource"
	tarantool "github.com/tarantool/go-tarantool"
)

// RecipesSvcInterface is an interface to test RecipesSvc
type RecipesSvcInterface interface {
	GetOne(recipeID int) (*resource.Recipe, error)
}

// RecipesSvc provides api to user end point
type RecipesSvc struct {
	Rsc resource.RecipesRscInterface
}

// NewRecipesSvc initiates RecipesSvc
func NewRecipesSvc(client *tarantool.Connection) *RecipesSvc {
	return &RecipesSvc{
		Rsc: resource.NewRecipesRsc(client),
	}
}

// GetOne gets user from users resouce
func (u *RecipesSvc) GetOne(recipesID int) (*resource.Recipe, error) {
	return u.Rsc.GetOne(recipesID)
}
