package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Time struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getTime)

	fmt.Println("Server is running at 127.0.0.1:3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := []Time{
		{CurrentTime: http.TimeFormat},
	}

	json.NewEncoder(w).Encode(currentTime)
}
