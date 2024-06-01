package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("String is collection of utf-8 character.")
	str1 := "hello world"
	fmt.Println(str1)

	fmt.Println("==============================")
	fmt.Println("Loop through the characters in string:")
	for index, val := range str1 {
		fmt.Println(index, val)
	}

	fmt.Println("================================")
	fmt.Println("Loop through the characters and display characters:")
	for index, val := range str1 {
		fmt.Println(index, string(val))
	}

	fmt.Println("=================================")
	fmt.Println("Split a string based on delimeterr:")
	str3 := "hello world"
	list_of_str := strings.Split(str3, " ")
	fmt.Println(list_of_str)
}
