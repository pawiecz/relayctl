package main

import "github.com/nathan-osman/go-rpigpio"

type PinMap map[string]int

type PinState struct {
	state bool
	value rpi.Value
}

const (
	PIN_ON = iota
	PIN_OFF
)

var modePinState = map[int]PinState{
	PIN_ON:  PinState{true, rpi.LOW},
	PIN_OFF: PinState{false, rpi.HIGH},
}

func AllPins() PinMap {
	pinmap := PinMap{
		"I1": 27,
		"I2": 17,
	}
	return pinmap
}

func SetupModule(pinmap PinMap) {
	for id, pin := range pinmap {
		p, err := rpi.OpenPin(pin, rpi.OUT)
		if err != nil {
			panic(err)
		}

		pins[id] = p
		module[id] = &Relay{ID: id, Pin: pin, State: false}
	}
}

func CloseModule() {
	for _, pin := range pins {
		pin.Close()
	}
}

var pins = make(map[string]*rpi.Pin)
