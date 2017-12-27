package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "relayctl reporting for duty!\n")
}

// Relay represents a single switch on controlled module.
type Relay struct {
	ID    string `json:"id"`
	Pin   int    `json:"pin"`
	State bool   `json:"state"`
}

type JsonResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

// module stores state of relays with their IDs as keys.
var module = make(map[string]*Relay)

// RelayIndex handles the relays index action (GET /relays).
func RelayIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var relays []*Relay
	for _, relay := range module {
		relays = append(relays, relay)
	}
	response := &JsonResponse{Data: &relays}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/relays", RelayIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}
