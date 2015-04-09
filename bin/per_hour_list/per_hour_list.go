package main

import (
	"io"
	"os"

	"flag"


	"github.com/bborbe/log"
	per_hour_entry "github.com/bborbe/stats/entry"
	per_hour_storage "github.com/bborbe/stats/storage"
	"github.com/bborbe/stats"
	io_util "github.com/bborbe/io/util"
	"fmt"
	"sort"
	"time"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL = "loglevel"
	PARAMETER_DB_PATH  = "db"
)

func main() {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	dbPathPtr := flag.String(PARAMETER_DB_PATH, stats.DEFAULT_DB_PATH, "path to database file")
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	writer := os.Stdout
	err := do(writer, *dbPathPtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, dbPath string) error {
	var err error
	dbPath, err = io_util.NormalizePath(dbPath)
	if err != nil {
		return err
	}
	storage := per_hour_storage.New(dbPath, false)
	var entries []per_hour_entry.Entry
	if entries, err = storage.FindEntrys(); err != nil {
		return err
	}
	sort.Sort(per_hour_entry.EntryByTimestamp(entries))

	for i := 0; i < len(entries)-1; i++ {
		a := entries[i]
		b := entries[i + 1]
		hours := float64(a.Timestamp - b.Timestamp) / float64(time.Hour)
		diff := float64(a.Value - b.Value) / hours
		d := time.Unix(0, a.Timestamp)
		fmt.Fprintf(writer, "%s %s\n", d.Format("2006-01-02 15:04:05"), extend(fmt.Sprintf("%.2f", diff), 12))
	}

	return nil
}

func extend(value string, length int) string {
	if (len(value) < length) {
		return extend(" "+value, length)
	}
	return value
}
