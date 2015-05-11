package list

import (
	stats_entry_service "github.com/bborbe/stats/entry/service"

	"net/http"

	"github.com/bborbe/log"
	error_handler "github.com/bborbe/server/handler/error"
	json_handler "github.com/bborbe/server/handler/json"
)

var (
	logger = log.DefaultLogger
)

type handler struct {
	service stats_entry_service.Service
}

func New(service stats_entry_service.Service) *handler {
	h := new(handler)
	h.service = service
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	list, err := h.service.List()
	if err != nil {
		logger.Debug(err)
		e := error_handler.NewErrorMessage(http.StatusInternalServerError, err.Error())
		e.ServeHTTP(responseWriter, request)
		return
	}
	j := json_handler.NewJsonHandler(list)
	j.ServeHTTP(responseWriter, request)
}
