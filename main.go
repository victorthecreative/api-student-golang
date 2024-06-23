package main

import (
	"log"

	"github.com/victorthecreative/api-student-golang/api"
)

func main() {

	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
