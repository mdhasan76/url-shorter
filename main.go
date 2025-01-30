package main

import (
	"fmt"
	"log"
	"net/http"
	route "url-shorter/router"
)

func main() {
	fmt.Println("Welcome to URL shorter backend. -developed by go lang")
	r := route.Route()

	fmt.Println("Hey server is running...")
	fmt.Println("Listening server is on 4001 port")
	err := http.ListenAndServe(":4001", r)
	fmt.Println(err, "this is err")
	log.Fatal()
}
