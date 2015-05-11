package service

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestNewHandlerImplementsService(t *testing.T) {
	r := New(nil)
	var i (*Service) = nil
	err := AssertThat(r, Implements(i).Message("check implements server.Server"))
	if err != nil {
		t.Fatal(err)
	}
}
