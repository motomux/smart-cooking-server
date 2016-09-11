package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/motomux/smart-cooking-server/controller"
)

func registerStatusHealthz(mux *httprouter.Router) {
	ctrl := &controller.StatusHealthzCtrl{}

	mux.GET("/_status/healthz", withGetCtrl(ctrl))
}
