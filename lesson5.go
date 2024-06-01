package main

import (
	"fmt"
	"strings"
)

func count_substring_k_charcter(str2 string, K int) int {
	var result int
	str3 := strings.ToLower(str2)
	runes := []rune(str3)
	list_of_alphabets := []rune("abcdefghijklmnopqrstuvwxyz")
	count := 0
	var arr [26]int
	for index, _ := range str3 {

		//Reset the memory.
		for index, _ := range arr {
			arr[index] = 0
		}

		count = 0

		//Calculate number of distinct elements.
		for j := index; j < len(str3); j++ {
			if arr[runes[j]-list_of_alphabets[0]] == 0 {
				count += 1
			}
			arr[runes[j]-list_of_alphabets[0]] += 1
			if count == K {
				result += 1
			}
			if count > K {
				break
			}
		}

	}
	return result
}

func main() {
	fmt.Println("Count the number substring having k distinct characters.")
	str1 := "abcbaa"
	result := count_substring_k_charcter(str1, 3)
	fmt.Println(result)
}
