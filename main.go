package main

import (
	"fmt"
	"log"
	server "server/dep"
)

func main() {
	srv := server.NewServer(":6969")
	go func() {
		for mssg := range srv.Msgch {
			fmt.Printf("%s: %s", mssg.From, mssg.Msg)
		}
	}()
	log.Fatal(srv.Start())
}
