package list

import (
	stats_entry_service "github.com/bborbe/stats/entry/service"

	"net/http"
)

type handler struct {
	service stats_entry_service.Service
}

func New(service stats_entry_service.Service) *handler {
	h := new(handler)
	h.service = service
	return h
}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

}
