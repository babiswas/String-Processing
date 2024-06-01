package main

import (
	"fmt"
)

func is_all_distinct_chars(str1 string) bool {
	list_of_runes := []rune(str1)
	mp := make(map[rune]int, len(str1))
	for _, val := range list_of_runes {
		_, ok := mp[val]
		if ok == false {
			mp[val] = 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Find distinct characters in string.")
	str1 := "hello world"
	mp := make(map[string]int, len(str1))
	for _, value := range str1 {
		temp_str := string(value)
		_, ok := mp[temp_str]
		if ok == false {
			mp[temp_str] = 1
		} else {
			mp[temp_str] = mp[temp_str] + 1
		}
	}
	fmt.Println(mp)
	test := is_all_distinct_chars(str1)
	fmt.Println(test)

}
