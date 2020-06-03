package main

import (
	"fmt"
)

func addHundred(num *int) {
	*num += 100
}

func main() {
	x := "here is an example of an address."
	// the & prefix calls the hexadecimal address of the variable.
	fmt.Println(&x)

	someNum := 75

	// user the * pointer to the address of someNum, it updates the value of the variable.
	addHundred(&someNum)

	// will print out 175
	fmt.Println(someNum)

	// creates a pointer to the address of someNum
	var someOtherNum *int = &someNum

	addHundred(someOtherNum)

	fmt.Println(someNum)

}
