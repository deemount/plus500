package main

import (
	"log"

	"github.com/deemount/plus500/api"
)

// Plus500 REST API
func main() {

	// assign error
	var err error

	// run application interface
	if err = api.Run(); err != nil {
		log.Fatalf("Plus500 REST API Error %s", err)
	}

}
