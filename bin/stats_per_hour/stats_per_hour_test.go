package main

import (
	"testing"

	"time"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
	stats_entry "github.com/bborbe/stats/entry"
)

func TestDo(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
	err = do(writer, "/tmp/per_hour_test.db")
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrintEntries(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
	entries := []stats_entry.Entry{stats_entry.Entry{Id: 1, Timestamp: 0, Value: 100}, stats_entry.Entry{Id: 1, Timestamp: 0 + int64(time.Hour), Value: 200}}
	err = printEntries(writer, entries)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(writer.Content(), NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(string(writer.Content()), Is("1970-01-01 02:00:00          200          100       100.00/h\n")); err != nil {
		t.Fatal(err)
	}

}
