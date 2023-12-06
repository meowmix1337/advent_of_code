package controller

import (
	"github.com/go-chi/chi/v5"

	log "github.com/sirupsen/logrus"
)

const APIV1 = "v1"

type Controller struct {
	Log    *log.Logger
	Router *chi.Mux
}

func NewController(log *log.Logger) *Controller {
	ctrl := new(Controller)
	ctrl.Log = log

	return ctrl
}
