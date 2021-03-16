package workout

import (
	"encoding/json"
	"fmt"
	"github.com/bradsk88/recroom/internal/api"
	"net/http"
)

type Handler struct {
}

type Response struct {
	Exercises []api.Exercise `json:"exercises"`
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	responseBody, err := json.Marshal(Response{
		Exercises: []api.Exercise{
			{
				ExerciseID:  "1",
				Name:        "Push Up",
				Repetitions: 10,
				Sets:        1,
			},
		},
	})

	if err != nil {
		writer.WriteHeader(500)
		_, err := writer.Write([]byte(fmt.Sprintf("Failed to generate workout: %s", err.Error())))
		if err != nil {
			fmt.Printf("Failed to write bytes: %s\n", err.Error())
		}
		return
	}

	_, err = writer.Write(responseBody)
	if err != nil {
		writer.WriteHeader(500)
		fmt.Printf("Failed to write bytes: %s\n", err.Error())
		return
	}
}

func NewHandler() *Handler {
	return &Handler{}
}
