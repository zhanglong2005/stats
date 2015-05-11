package main

import (
	"flag"
	"net/http"
	"os"

	"io"

	io_util "github.com/bborbe/io/util"
	"github.com/bborbe/log"
	"github.com/bborbe/stats"
	"github.com/bborbe/stats/handler"
	"github.com/facebookgo/grace/gracehttp"
)

const (
	PARAMETER_LOGLEVEL = "loglevel"
	PARAMETER_DB_PATH  = "db"
)

var (
	logger          = log.DefaultLogger
	addressPtr      = flag.String("a0", ":48568", "Zero address to bind to.")
	documentRootPtr = flag.String("root", "", "Document root directory")
	logLevelPtr     = flag.String("loglevel", log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	dbPathPtr       = flag.String(PARAMETER_DB_PATH, stats.DEFAULT_DB_PATH, "path to database file")
)

func main() {
	defer logger.Close()
	flag.Parse()

	writer := os.Stdout
	err := do(writer, *addressPtr, *documentRootPtr, *dbPathPtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, address string, documentRoot string, dbPath string) error {
	dbPath, err := io_util.NormalizePath(dbPath)
	if err != nil {
		return err
	}
	gracehttp.Serve(createServer(address, documentRoot, dbPath))
	return nil
}

func createServer(address string, documentRoot string, dbPath string) *http.Server {
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)
	return &http.Server{Addr: address, Handler: handler.NewHandler(documentRoot, dbPath)}
}
