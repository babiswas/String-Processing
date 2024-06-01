package main

import (
	"fmt"
	"strings"
)

func find_all_substring_string(str1 string) []string {
	list_of_substring := make([]string, len(str1))
	temp := ""
	for index, value := range str1 {
		temp += string(value)
		for j := index + 1; j < len(str1); j++ {
			temp += string(str1[j])
			list_of_substring = append(list_of_substring, strings.Clone(temp))
		}
		temp = ""
	}
	return list_of_substring
}

func main() {
	fmt.Println("Find the list of all substring of a string:")
	str1 := "hello world"
	list_of_substring := find_all_substring_string(str1)
	fmt.Println(list_of_substring)
}
