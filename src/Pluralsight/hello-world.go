package main

// This is a simple "hello world" type program


// import other packages we need
import (
    "fmt"
    "runtime"
)


// main function of the main package is our starting point
func main() {

    fmt.Println("Hello from " ,  runtime.GOOS)
 
    
}