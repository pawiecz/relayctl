package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/relays", RelayIndex)
	router.GET("/relays/:id", RelayShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
