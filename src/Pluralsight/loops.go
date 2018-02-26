package main

import (
	"fmt"
)

func main() {

	// the standard mechanism for loops in go is the for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for is used as a while
	x := 0
	for x < 1000 {
		x += 1
	}
	fmt.Println(x)

	// and is used for an infinite loop
	// for {
	// }

}
