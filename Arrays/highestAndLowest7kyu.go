package main

import (
	"fmt"
	"sort"
	s "strings"
)

func main() {

	str1 := "1 2 3 7 -3"
	str2 := "13 18 8 44 2 104 -18"

	arr1 := s.Fields(str1)
	arr2 := s.Fields(str2)

	sort.Strings(arr1)
	fmt.Println(arr1, arr2)

}
