package main

import "fmt"

var myText = "Hello, World"
var pText = &myText

func main() {

	aString := "INITIAL"
	changeIt(&aString)

	fmt.Println(aString)
}

func changeIt(inVal *string) string {
	*inVal = "CHANGED"
	return *inVal
}
