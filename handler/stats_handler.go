package handler

import (
	"net/http"

	"github.com/bborbe/log"
	"github.com/bborbe/server/handler/cachingheader"
	"github.com/bborbe/server/handler/contenttype"
	"github.com/bborbe/server/handler/fallback"
	log_handler "github.com/bborbe/server/handler/log"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder"
	"github.com/bborbe/server/handler_finder/part"
	"github.com/bborbe/server/handler_finder/rest"
	stats_entry_handler_create "github.com/bborbe/stats/entry/handler/create"
	stats_entry_handler_list "github.com/bborbe/stats/entry/handler/list"
	stats_entry_service "github.com/bborbe/stats/entry/service"
	stats_entry_storage "github.com/bborbe/stats/entry/storage"
)

var logger = log.DefaultLogger

func NewHandler(documentRoot string, dbPath string) http.Handler {
	logger.Debugf("root: %s", documentRoot)

	entryStorage := stats_entry_storage.New(dbPath, false)
	entryService := stats_entry_service.New(entryStorage)

	fileServer := cachingheader.NewCachingHeaderHandler(contenttype.NewContentTypeHandler(http.FileServer(http.Dir(documentRoot))))
	handlerFinder := part.New("")
	handlerFinder.RegisterHandler("/", fileServer)
	handlerFinder.RegisterHandler("/css", fileServer)
	handlerFinder.RegisterHandler("/js", fileServer)
	handlerFinder.RegisterHandler("/images", fileServer)
	handlerFinder.RegisterHandlerFinder("/entry", createEntryHandlerFinder("/entry", entryService))
	return log_handler.NewLogHandler(fallback.NewFallback(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", 404)))
}

func createEntryHandlerFinder(prefix string, entryService stats_entry_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(stats_entry_handler_list.New(entryService))
	hf.RegisterCreateHandler(stats_entry_handler_create.New(entryService))
	return hf
}
