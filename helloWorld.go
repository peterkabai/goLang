// This is a comment

package main

import "fmt"

const text string = "This is constant text"

func main() {
	var x string
	x = "Hello, World"
	fmt.Println(x)
	
	y := "Hello, Again"
	fmt.Println(y)
	
	next()
}

func next() {
	fmt.Println("My name is Peter")
}