package main

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestDo(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
	err = do(writer, "/tmp/per_hour_test.db", "1337")
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
}
