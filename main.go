package main

import (
	"log"
	"personal-portfolio/server"
)

func main() {
	server := server.NewServer()
	log.Fatal(server.Start())
}
