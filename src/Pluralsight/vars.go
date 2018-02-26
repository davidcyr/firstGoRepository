package main

import (
	"fmt"
)

var (
	x, y     int
	myString = "implied"
)

func main() {

	// if this isn't commented out, it will give an error if the
	// name is declared but not used
	//var z bool

	fmt.Println(myString)

}
