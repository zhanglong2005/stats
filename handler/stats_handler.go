package handler

import (
	"net/http"

	"github.com/bborbe/log"
	"github.com/bborbe/server/handler/cachingheader"
	"github.com/bborbe/server/handler/contenttype"
	"github.com/bborbe/server/handler/fallback"
	log_handler "github.com/bborbe/server/handler/log"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder/part"
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
	handlerFinder.RegisterHandler("/stats/insert", stats_handler_insert.New())
	handlerFinder.RegisterHandler("/stats/list", stats_handler_list.New())
	return log_handler.NewLogHandler(fallback.NewFallback(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", 404)))
}
