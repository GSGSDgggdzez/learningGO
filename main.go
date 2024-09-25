package main

import "fmt"

func main() {
	fmt.Println("hello Bro these is  my first go program")
	// these is how we declare an integer in Go
	var integer, string = 1, "hello"

	fmt.Println(integer, string)
	// short cut
	booleans := true

	fmt.Println(booleans)
	// pointer
	x := 1
	P := &x

	// Problem: cannot use P (variable of type *int) as []byte value in argument to fmt.Appendln
	fmt.Println(*P)
	// type

	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	fmt.Println(c, f)
}
