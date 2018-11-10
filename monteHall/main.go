// In the Monty Hall problem, a contestant has three doors,
// two with goats behind it and one with a car. After they select a door,
// the host opens one of the doors that was not picked.
// the contestant then has the opportunity to keep their selected door, or to switch.

package main

import "fmt"
import "math/rand"
import "time"

func main() {

	repeats := 10000

	count := 0
	i := 0
	for i < repeats {
		doors := createDoors()                    // create a set of doors
		ourDoor := random(0, 2)                   // picks a random door
		doorToOpen := randomButNot(0, 2, ourDoor) // picks a different door
		if doors[doorToOpen] != 1 {               // if the different randomly opened door is not a car
			if doors[ourDoor] == 1 {
				count++
			}
		}
		i++
	}
	fmt.Print(float64(count) / float64(repeats))
	fmt.Println(" probability when not switched")

	count = 0
	i = 0
	for i < repeats {
		doors := createDoors()                    // create a set of doors
		ourDoor := random(0, 2)                   // picks a random door
		doorToOpen := randomButNot(0, 2, ourDoor) // picks a different door
		if doors[doorToOpen] == 1 {               // if the different randomly opened door is a car

			// switches the door, so that when opened it does not have a car
			if doorToOpen == 0 && ourDoor == 1 {
				doorToOpen = 2
			} else if doorToOpen == 0 && ourDoor == 2 {
				doorToOpen = 1
			} else if doorToOpen == 1 && ourDoor == 0 {
				doorToOpen = 2
			} else if doorToOpen == 1 && ourDoor == 2 {
				doorToOpen = 0
			} else if doorToOpen == 2 && ourDoor == 0 {
				doorToOpen = 1
			} else if doorToOpen == 2 && ourDoor == 1 {
				doorToOpen = 0
			}
		}

		// switches the door originally selected by the user
		if doorToOpen == 0 && ourDoor == 1 {
			ourDoor = 2
		} else if doorToOpen == 0 && ourDoor == 2 {
			ourDoor = 1
		} else if doorToOpen == 1 && ourDoor == 0 {
			ourDoor = 2
		} else if doorToOpen == 1 && ourDoor == 2 {
			ourDoor = 0
		} else if doorToOpen == 2 && ourDoor == 0 {
			ourDoor = 1
		} else if doorToOpen == 2 && ourDoor == 1 {
			ourDoor = 0
		}

		// if the door selected does have a car
		if doors[ourDoor] == 1 {
			count++
		}

		i++
	}
	fmt.Print(float64(count) / float64(repeats))
	fmt.Println(" probability when switched")

}

// generates a random number, low to high inclusive
func random(low int, high int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(high-low+1) + low
}

// generates a random number, low to high inclusive
func randomButNot(low int, high int, not int) int {
	rand.Seed(time.Now().UnixNano())

	newRandom := not
	for newRandom == not {
		newRandom = rand.Intn(high-low+1) + low
	}
	return newRandom
}

// creates
func createDoors() [3]int {

	// get a random number between 0 and 2
	randomInt := random(0, 2)

	// create an array, length of 3 for the doors
	var doors [3]int

	// add a 1 or 0 to each door
	i := 0
	for i < 3 {
		if i == randomInt {
			doors[i] = 1 // randomly add the one
		} else {
			doors[i] = 0 // by default add a zero
		}
		i++
	}
	return doors
}
