package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str1 string) bool {
	temp := ""
	for i := len(str1) - 1; i >= 0; i-- {
		temp += string(str1[i])
	}
	if strings.Compare(temp, str1) == 0 {
		return true
	} else {
		return false
	}
}

func find_all_substrings(str1 string) []string {
	l := make([]string, len(str1))
	temp := ""
	for index, _ := range str1 {
		for i := index; i < len(str1); i++ {
			temp += string(str1[i])
			l = append(l, strings.Clone(temp))
		}
		temp = ""
	}
	return l
}

func main() {
	fmt.Println("Find all substrings of string:")
	str1 := "forgeeksskeegfor"
	all_substring := find_all_substrings(str1)
	max_length := -1
	var max_str string
	for _, val := range all_substring {
		if isPalindrome(val) {
			if len(val) > max_length {
				max_length = len(val)
				max_str = val
			}
		}
	}
	fmt.Println("The maximum palindromic substring is:", max_str)
	fmt.Println("The maximum length of palindromic substring is:", max_length)

}
