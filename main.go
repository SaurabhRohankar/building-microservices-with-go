package main

import (
	"github.com/SaurabhRohankar/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "rpoduct-api", log.LstdFlags)

	//Create the handler
	hh := h.NewHello(l)

	//create new ServeMux
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// Listen for connection an all ip address (0.0.0.0)
	// port 8080
	log.Println("Starting Server")
	err := http.ListenAndServe(":8080", sm)
	log.Fatal(err)
}
