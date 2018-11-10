package main

import "fmt"
import "math/rand"
import "time"

func main() {

	numbersToPrint := 50

	i := 0
	for i < numbersToPrint {

		// randomizes the seed
		rand.Seed(time.Now().UnixNano())

		// random number between high and low
		randomInt := randomButNot(1, 3, 2)
		fmt.Println(randomInt)

		i++
	}
}

func randomButNot(low int, high int, not int) int {
	rand.Seed(time.Now().UnixNano())

	newRandom := not
	for newRandom == not {
		newRandom = rand.Intn(high-low+1) + low
	}
	return newRandom
}
