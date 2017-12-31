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
		writeError(w, err)
		return
	}
	writeResponse(w, http.StatusOK, &JsonResponse{Data: relay})
}

// RelayOn handles the relays on action (GET /relays/:id/on).
func RelayOn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	switchRelay(w, id, PIN_ON)
}

// RelayOff handles the relays off action (GET /relays/:id/off).
func RelayOff(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	switchRelay(w, id, PIN_OFF)
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
func switchRelay(w http.ResponseWriter, id string, mode int) {
	relay, err := findRelay(id)
	if err != nil {
		writeError(w, err)
		return
	}
	if err := setState(id, mode); err != nil {
		writeError(w, err)
		return
	}
	writeResponse(w, http.StatusOK, &JsonResponse{Data: relay})
}

// setState switches given pin to requested state.
func setState(id string, mode int) error {
	ps := modePinState[mode]
	if err := pins[id].Write(ps.value); err != nil {
		// Writing requested state to the given pin failed
		return ErrStateChangeFailed
	}
	module[id].State = ps.state
	return nil
}

// writeError wraps writing API error response.
func writeError(w http.ResponseWriter, err error) {
	status := errHttpStatus[err]
	apiError := &ApiError{Status: status, Title: err.Error()}
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
