package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		loop(i)
		i++
	}
}

func loop(i int) {
	fmt.Println(i)
}