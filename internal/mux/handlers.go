package mux

import (
	"github.com/bradsk88/recroom/internal/api/workout"
	"net/http"
)

type Handlers struct {
	mux *http.ServeMux
}

func (h Handlers) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.mux.ServeHTTP(writer, request)
}

func New() *Handlers {
	mux := http.NewServeMux()
	mux.Handle("/workout/generate", workout.NewHandler())

	return &Handlers{
		mux: mux,
	}
}
