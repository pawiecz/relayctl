package main

import (
	"errors"
	"net/http"
)

var (
	// ErrRelayNotFound is returned when Relay is not found.
	ErrRelayNotFound = errors.New("relay not found")
	// ErrStateChangeFailed is returned if state change fails.
	ErrStateChangeFailed = errors.New("state change failed")
)

var errHttpStatus = map[error]int{
	ErrRelayNotFound:     http.StatusNotFound,
	ErrStateChangeFailed: http.StatusServiceUnavailable,
}
