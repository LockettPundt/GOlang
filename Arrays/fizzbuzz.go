package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr [101]string

	for i := 1; i < len(arr); i++ {
		if (i%3 == 0) && (i%5 == 0) {
			arr[i] = "fizzbuzz"
		} else if i%3 == 0 {
			arr[i] = "fizz"
		} else if i%5 == 0 {
			arr[i] = "buzz"
		} else {
			arr[i] = strconv.Itoa(i)
		}
	}

	fmt.Println(arr)

}
