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
	writeResponse(w, http.StatusOK, &JsonResponse{Data: relays})
}

// RelayShow handles the relays show action (GET /relays/:id).
func RelayShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	relay, err := findRelay(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeResponse(w, http.StatusOK, &JsonResponse{Data: relay})
}

// findRelay locates Relay with the given ID on the module.
func findRelay(id string) (*Relay, error) {
	relay, ok := module[id]
	if !ok {
		// No relay with the ID given in the URL has been found
		return nil, ErrRelayNotFound
	}
	return relay, nil
}

// writeError wraps writing API error response.
func writeError(w http.ResponseWriter, status int, title string) {
	apiError := &ApiError{Status: status, Title: title}
	writeResponse(w, status, &JsonErrorResponse{Error: apiError})
}

// writeResponse writes standard JSON API response with status code.
func writeResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		apiError := &ApiError{Status: http.StatusInternalServerError, Title: "internal server error"}
		writeResponse(w, http.StatusInternalServerError, apiError)
	}
}
