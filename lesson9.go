package main

import (
	"fmt"
	"strings"
)

func is_palindrome(str1 string) bool {
	temp := ""
	for i := len(str1) - 1; i >= 0; i-- {
		temp += string(str1[i])
	}
	if strings.Compare(str1, temp) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Verify if the string is a palindrome")
	if is_palindrome("12321") {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
