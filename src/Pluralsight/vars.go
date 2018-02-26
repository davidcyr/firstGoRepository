package main

import (
	"fmt"
)

var (
	x, y     int
	myString = "implied"
)

const pi = 3.14

func main() {

	// if this isn't commented out, it will give an error if the
	// name is declared but not used
	//var z bool

	fmt.Println(myString)

	// you aren't forced to use constants
	const j = 2

}
