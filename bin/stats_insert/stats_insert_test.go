package main

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestDo(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
	err = do(writer, "/tmp/stats_test.db", "1337")
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
}
