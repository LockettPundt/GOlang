package main

import (
	"fmt"
)

// declaring argument type and return value type.
func sayHello(message, name string) string {
	return fmt.Sprintf("%s %s, how are you?", message, name)
}

// returning multiple values.
func sayHiTwice(message1, message2 string) (string, string) {
	firstMessageToSend := fmt.Sprintf("%s! This is the first message.", message1)
	secondMessageToSend := fmt.Sprintf("%s! This is the second message.", message2)
	return firstMessageToSend, secondMessageToSend
}

func main() {
	var greeting string = "Hi There"
	name := "Lockett"
	fmt.Println(sayHello(greeting, name))

	// assigning the 2 returned values from sayHiTwice to one and two.
	var one, two string = sayHiTwice(greeting, "Salutations")
	fmt.Println(one, two)
}
