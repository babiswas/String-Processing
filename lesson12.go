package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReverseString(str1 string) string {
	runes := []rune(str1)
	index1 := 0
	index2 := len(str1) - 1

	for index1 < index2 {
		runes[index1], runes[index2] = runes[index2], runes[index1]
		index1 += 1
		index2 -= 1
	}

	return string(runes)
}

func Reverse(s string) string {
	runes := []rune(s)
	index := len(s) - 1
	start_index := 0
	test := make([]rune, len(s))
	for index >= 0 {
		test[start_index] = runes[index]
		index -= 1
		start_index += 1
	}
	return string(test)
}

func main() {
	fmt.Println("Attempt to reverse the string.")
	fmt.Println("Enter the string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 := scanner.Text()
	reversed_string := ReverseString(str1)
	fmt.Println(reversed_string)
	fmt.Println(Reverse(reversed_string))
}
