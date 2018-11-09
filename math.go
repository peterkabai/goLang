package main

import "fmt"

func main() {
	var first int
	fmt.Println("Input first number:")
	fmt.Scanf("%d", &first)

	var second int
	fmt.Println("Input second number:")
	fmt.Scanf("%d", &second)
	
	v := division(first, second)
	fmt.Println(v)
}

func addition(x int, y int) (int) {
	var sum = x + y
	return(sum)
}

func subtraction(x int, y int) (int) {
	var sub = x + y
	return(sub)
}

func multiplication(x int, y int) (int) {
	var multiple = x * y
	return(multiple)
}

func division(x int, y int) (int) {
	var div = x / y
	return(div)
}