package main

import "errors"

var (
	// ErrRelayNotFound is returned when Relay is not found.
	ErrRelayNotFound = errors.New("relay not found")
	// ErrStateChangeFailed is returned if state change fails.
	ErrStateChangeFailed = errors.New("state change failed")
)
