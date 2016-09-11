package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/motomux/smart-cooking-server/controller"
)

func registerRecipes(mux *httprouter.Router, env *Env) {
	ctrl := controller.NewRecipesCtrl(env.Client)
	mux.GET("/recipes/:id", withGetOneCtrl(ctrl))
}
