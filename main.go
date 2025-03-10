package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shorter/controller"
	route "url-shorter/router"
)

func main() {
	controller.ServerMainFn()
	fmt.Println("Welcome to URL shorter backend. -developed by go lang")
	r := route.Route()

	fmt.Println("Hey server is running...")
	fmt.Println("Listening server is on 4001 port")
	log.Fatal(http.ListenAndServe(":4001", r))
	// log.Fatal()
}
