package main

// Relay represents a single switch on controlled module.
type Relay struct {
	ID    string `json:"id"`
	Pin   int    `json:"pin"`
	State bool   `json:"state"`
}

// module stores state of relays with their IDs as keys.
var module = make(map[string]*Relay)
