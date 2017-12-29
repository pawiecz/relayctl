package main

import (
	"log"
	"net/http"
)

func main() {
	SetupModule(AllPins())
	defer CloseModule()

	router := NewRouter(AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}
