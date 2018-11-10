package main

import "fmt"
import "math/rand"

func main() {

	low := 1
	high := 10
	numbersToPrint := 10

	i := 0
	for i < numbersToPrint {

		// randomizes the seed
		rand.Seed(time.Now().UnixNano())

		// random number between high and low
		randomInt := rand.Intn(high-low+1) + low
		fmt.Println(randomInt)

		i++
	}
}
