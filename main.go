package main

import (
	"fmt"
	"url-shorter/controller"
)

func main() {
	fmt.Println("Welcome to URL shorter backend. -developed by go lang")
	a := controller.MoneDorse("hello")
	controller.UppercasingFuncion()
	fmt.Println(a)
}
