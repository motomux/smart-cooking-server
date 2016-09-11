package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/motomux/smart-cooking-server/controller"
	tarantool "github.com/tarantool/go-tarantool"
)

// Env is env values
type Env struct {
	Client *tarantool.Connection
}

// NewHandler inititializes mux and register handlers
func NewHandler(env *Env) *httprouter.Router {
	mux := httprouter.New()

	registerStatusHealthz(mux)

	registerRecipes(mux, env)

	return mux
}

func withGetOneCtrl(ctrl controller.GetOneCtrlInterface) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctrl.GetOne(w, r, ps)
	}
}

func withGetCtrl(ctrl controller.GetCtrlInterface) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctrl.Get(w, r, ps)
	}
}

func withPostCtrl(ctrl controller.PostCtrlInterface) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctrl.Post(w, r, ps)
	}
}
