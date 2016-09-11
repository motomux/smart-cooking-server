package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetOneCtrlInterface provides GetOne method
type GetOneCtrlInterface interface {
	GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// GetCtrlInterface provides Get method
type GetCtrlInterface interface {
	Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// PostCtrlInterface provides Post method
type PostCtrlInterface interface {
	Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// PutCtrlInterface provides Put method
type PutCtrlInterface interface {
	Put(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// PatchCtrlInterface provides Patch method
type PatchCtrlInterface interface {
	Patch(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// DeleteCtrlInterface provides Delete method
type DeleteCtrlInterface interface {
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
