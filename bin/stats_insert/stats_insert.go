package main

import (
	"io"
	"os"

	"flag"

	"strconv"
	"time"

	io_util "github.com/bborbe/io/util"
	"github.com/bborbe/log"
	"github.com/bborbe/stats"
	stats_entry "github.com/bborbe/stats/entry"
	stats_entry_storage "github.com/bborbe/stats/entry/storage"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL = "loglevel"
	PARAMETER_VALUE    = "value"
	PARAMETER_DB_PATH  = "db"
)

var (
	logLevelPtr = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	valuePtr    = flag.String(PARAMETER_VALUE, "", "value")
	dbPathPtr   = flag.String(PARAMETER_DB_PATH, stats.DEFAULT_DB_PATH, "path to database file")
)

func main() {
	defer logger.Close()
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	writer := os.Stdout
	err := do(writer, *dbPathPtr, *valuePtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, dbPath string, valueString string) error {
	var err error
	dbPath, err = io_util.NormalizePath(dbPath)
	if err != nil {
		return err
	}
	var value int
	if value, err = strconv.Atoi(valueString); err != nil {
		return err
	}
	storage := stats_entry_storage.New(dbPath, false)
	timestamp := time.Now().UnixNano()
	entry := &stats_entry.Entry{Value: value, Timestamp: timestamp}
	return storage.CreateEntry(entry)
}
