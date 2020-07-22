package inquiry

import (
	"github.com/gorilla/mux"
	redmpay "tapera.integrasi/service/redempay"
)

// Controller struct
type Controller struct {
	srv redmpay.Service
}

// NewController func
func NewController(srv redmpay.Service) *Controller {
	return &Controller{srv: srv}
}

// Route func
func (c *Controller) Route(r *mux.Router) {
	//routes grouping
	s := r.PathPrefix("/dummy/bri").Subrouter()
	s.HandleFunc("/test", c.Test).Methods("GET")
	s.HandleFunc("/redempay", c.RedemPay).Methods("POST")
}
