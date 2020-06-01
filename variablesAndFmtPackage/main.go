package main

import (
	"fmt"
	t "time"
)

// the 't' is renaming the import to whatever you like. like 'as'

func main() {

	var helloWorld string = "Hello World"
	helloWorldAgain := "Hello Again, World"      // := declares the variable type and value.
	printThis := fmt.Sprintln("Print this.....") // Sprint returns a value instead of prints.

	fmt.Println(helloWorld, helloWorldAgain, printThis)
	fmt.Printf("This is how you interpolate strings ----> %v, %v, %v", helloWorld, helloWorldAgain, printThis)
	// string interpolation. can be done with any of the print methods.

	var greeting string
	fmt.Scan(&greeting) // user input prompt and stored in the greeting variable.

	fmt.Printf("%v, it is now %v", greeting, t.Now()) // prints the current time from the tim package

}
