package main

import (
	"fmt"
	"strings"
)

func is_first_rotation_second(str1, str2 string) bool {
	str3 := str2 + str2
	if strings.Contains(str3, str1) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Verify if the first string is rotation of the second string")
	str1 := "abc"
	str2 := "bca"
	if is_first_rotation_second(str2, str1) {
		fmt.Printf("%s is the rotation of second string %s \n", str2, str1)
	} else {
		fmt.Printf("%s is not the rotation of second string %s \n", str2, str1)
	}
}
