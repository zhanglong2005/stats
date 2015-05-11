package main

import (
	"io"
	"os"

	"flag"

	"fmt"
	"sort"
	"time"

	io_util "github.com/bborbe/io/util"
	"github.com/bborbe/log"
	"github.com/bborbe/stats/per_hour"
	per_hour_entry "github.com/bborbe/stats/per_hour/entry"
	per_hour_storage "github.com/bborbe/stats/per_hour/storage"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL = "loglevel"
	PARAMETER_DB_PATH  = "db"
)

func main() {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	dbPathPtr := flag.String(PARAMETER_DB_PATH, per_hour.DEFAULT_DB_PATH, "path to database file")
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
	var entries []per_hour_entry.Entry
	if entries, err = readEntriesFromDb(dbPath); err != nil {
		return err
	}
	return printEntries(writer, entries)
}

func readEntriesFromDb(dbPath string) ([]per_hour_entry.Entry, error) {
	var err error
	var entries []per_hour_entry.Entry
	dbPath, err = io_util.NormalizePath(dbPath)
	if err != nil {
		return nil, err
	}
	storage := per_hour_storage.New(dbPath, false)
	if entries, err = storage.FindEntrys(); err != nil {
		return nil, err
	}
	return entries, nil
}

func printEntries(writer io.Writer, entries []per_hour_entry.Entry) error {
	sort.Sort(per_hour_entry.EntryByTimestamp(entries))

	for i := 0; i < len(entries)-1; i++ {
		a := entries[i]
		b := entries[i+1]
		hourDiff := float64(a.Timestamp-b.Timestamp) / float64(time.Hour)
		valueDiff := a.Value - b.Value
		diff := float64(valueDiff) / hourDiff
		t := time.Unix(0, a.Timestamp)
		fmt.Fprintf(writer, "%s %s %s/h\n", timeToString(t), extendToLength(fmt.Sprintf("%d", valueDiff), 12), extendToLength(fmt.Sprintf("%.2f", diff), 12))
	}
	return nil
}

func timeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func extendToLength(value string, length int) string {
	if len(value) < length {
		return extendToLength(" "+value, length)
	}
	return value
}
