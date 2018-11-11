package main

import "fmt"
import "regexp"
import "strings"
import "strconv"

func main() {

	fmt.Println(getFloors("12u4d2d"))        // 6
	fmt.Println(getFloors("3u9u4d"))         // 8
	fmt.Println(getFloors("1u2u3d4u5d6u7d")) // -2

}

func getFloors(input string) int64 {

	// splits at every u or d
	change := strings.FieldsFunc(input, func(r rune) bool {
		switch r {
		case 'd', 'u':
			return true
		}
		return false
	})

	// splits at every digit
	reg := regexp.MustCompile(`\d+`)
	direction := reg.Split(input, -1)

	var floor int64 = 0

	// for each direction, adds or subtracts the floor
	i := 0
	for i < len(direction)-1 {

		s, err := strconv.ParseInt(change[i], 10, 0)

		if err == nil {
			if direction[i+1] == "u" {
				floor += s
			} else {
				floor -= s
			}
			i++
		} else {
			fmt.Println("There was an error")
		}
	}

	return floor
}
