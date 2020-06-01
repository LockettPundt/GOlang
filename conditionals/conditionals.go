package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// creates a new seed for each run. otherwise the randomizer will
	// repeat the same number each time.

	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(44)

	if number == 42 {
		fmt.Println("Yes, it's 42")
	} else {
		fmt.Println("No sir.")
	}

	// else if example as well as logical operators &&, !, and ||

	if num := 700; (number == num) && (number%2 != 0) {
		fmt.Println("Very good.")
	} else if num == 700 && num-number >= 600 {
		fmt.Println("just an else if here.")
	}

	// switch statement example. you

	switch someVariable := 88; someVariable {
	case 77:
		fmt.Println("Do this.")
	case 87:
		fmt.Println("Do that.")
	case 88:
		fmt.Println("Dance.")
	default:
		fmt.Println("None of the above.")
	}

}
