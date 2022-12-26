package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocantodev/api/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type ResponseError struct {
	Message string `json:"message"`
}

type Response struct {
	Configuration config.Config `json:"configuration"`
	Seed          string        `json:"seed"`
	SeedName      string        `json:"seedName"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/config", getConfig)

	fmt.Println("Server is running at 127.0.0.1:3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	configuration, err := config.Make()

	if err != nil {
		w.WriteHeader(500)
		response := ResponseError{Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Response{
		Configuration: *configuration,
		SeedName:      (*configuration).App.Env,
		Seed:          os.Getenv((*configuration).App.Env),
	}

	json.NewEncoder(w).Encode(response)
}
