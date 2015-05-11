package list

import "net/http"

type handler struct {
}

func New() *handler {
	return new(handler)
}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

}
