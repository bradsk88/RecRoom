package main

import (
	"fmt"
	"github.com/bradsk88/recroom/internal/mux"
	"net/http"
)

func main() {
	m := mux.New()
	err := http.ListenAndServe(":8100", m)
	if err != nil {
		fmt.Printf("Failed to server: %s", err.Error())
		return
	}
}
