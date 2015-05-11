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
	stats_handler_insert "github.com/bborbe/stats/handler/insert"
	stats_handler_list "github.com/bborbe/stats/handler/list"
)

var logger = log.DefaultLogger

func NewHandler(documentRoot string, dbPath string) http.Handler {
	logger.Debugf("root: %s", documentRoot)

	fileServer := cachingheader.NewCachingHeaderHandler(contenttype.NewContentTypeHandler(http.FileServer(http.Dir(documentRoot))))
	handlerFinder := part.New("")
	handlerFinder.RegisterHandler("/", fileServer)
	handlerFinder.RegisterHandler("/css", fileServer)
	handlerFinder.RegisterHandler("/js", fileServer)
	handlerFinder.RegisterHandler("/images", fileServer)
	handlerFinder.RegisterHandlerFinder("/entry", createEntryHandlerFinder("/entry"))
	return log_handler.NewLogHandler(fallback.NewFallback(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", 404)))
}

func createEntryHandlerFinder(prefix string) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(stats_handler_list.New())
	hf.RegisterCreateHandler(stats_handler_insert.New())
	return hf
}
