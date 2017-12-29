package main

import "github.com/nathan-osman/go-rpigpio"

type PinMap map[string]int

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
