package list

import (
	stats_entry "github.com/bborbe/stats/entry"
	stats_entry_service "github.com/bborbe/stats/entry/service"

	"net/http"

	"strconv"

	"sort"

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

type parameter struct {
	limit int
}

func New(service stats_entry_service.Service) *handler {
	h := new(handler)
	h.service = service
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	err := h.serveHTTP(responseWriter, request)
	if err != nil {
		logger.Debug(err)
		e := error_handler.NewErrorMessage(http.StatusInternalServerError, err.Error())
		e.ServeHTTP(responseWriter, request)
		return
	}
}

func (h *handler) serveHTTP(responseWriter http.ResponseWriter, request *http.Request) error {
	var err error
	var limit int
	if limit, err = strconv.Atoi(request.FormValue("limit")); err != nil {
		limit = -1
	}
	var list []stats_entry.Entry
	if limit > 0 {
		if list, err = h.service.ListLimited(limit); err != nil {
			return err
		}
	} else {
		if list, err = h.service.List(); err != nil {
			return err
		}
	}

	sort.Sort(stats_entry.EntryByTimestampDesc(list))

	j := json_handler.NewJsonHandler(list)
	j.ServeHTTP(responseWriter, request)
	return nil
}
