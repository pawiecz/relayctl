package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index handles the default route (GET /).
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "relayctl reporting for duty!\n")
}

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

// RelayShow handles the relays show action (GET /relays/:id).
func RelayShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	relay, ok := module[id]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !ok {
		// No relay with the ID given in the URL has been found
		w.WriteHeader(http.StatusNotFound)
		response := JsonErrorResponse{
			Error: &ApiError{
				Status: http.StatusNotFound,
				Title:  "Relay not found",
			},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
		return
	}
	response := JsonResponse{Data: relay}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
