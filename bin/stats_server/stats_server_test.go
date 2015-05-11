package main

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestCreateServer(t *testing.T) {
	server := createServer(":45678", "/tmp", "/tmp/stats_test.db")
	err := AssertThat(server, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}
