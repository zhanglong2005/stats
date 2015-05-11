package insert

import (
	"testing"

	"net/http"

	. "github.com/bborbe/assert"
)

func TestImplementsHttpHandler(t *testing.T) {
	r := New()
	var i *http.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
