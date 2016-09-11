package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// StatusHealthzCtrl is a controller for status check
type StatusHealthzCtrl struct{}

// NewStatusHealthzCtrl initializes StatusHealthzCtrl
func NewStatusHealthzCtrl() *StatusHealthzCtrl {
	return &StatusHealthzCtrl{}
}

// Get writes response with 204 status code
func (s *StatusHealthzCtrl) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	respond(w, r, http.StatusNoContent, nil)
}
