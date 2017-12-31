package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nathan-osman/go-rpigpio"
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

// RelayOn handles the relays on action (GET /relays/:id/on).
func RelayOn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	switchRelay(w, id, true, rpi.LOW)
}

// RelayOff handles the relays off action (GET /relays/:id/off).
func RelayOff(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	switchRelay(w, id, false, rpi.HIGH)
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

// switchRelay provides unified wrapper for Relay On/Off actions.
func switchRelay(w http.ResponseWriter, id string, state bool, pinState rpi.Value) {
	relay, err := findRelay(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := setState(id, state, pinState); err != nil {
		writeError(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	writeResponse(w, http.StatusOK, &JsonResponse{Data: relay})
}

// setState switches given pin to requested state.
func setState(id string, state bool, pinState rpi.Value) error {
	if err := pins[id].Write(pinState); err != nil {
		// Writing requested state to the given pin failed
		return ErrStateChangeFailed
	}
	module[id].State = state
	return nil
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
