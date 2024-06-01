package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func find_even_palindrome(rns []rune, start_index, end_index int, lngth int) string {
	for start_index >= 0 && end_index < lngth {
		if rns[start_index] == rns[end_index] {
			start_index -= 1
			end_index += 1
		} else {
			start_index += 1
			break
		}
	}
	return string(rns[start_index:end_index])
}

func find_odd_palindrome(rns []rune, start_index, end_index int, lngth int) string {
	for start_index >= 0 && end_index <= lngth {
		if rns[start_index] == rns[end_index] {
			start_index -= 1
			end_index += 1
		} else {
			start_index += 1
			break
		}
	}
	return string(rns[start_index:end_index])
}

func find_longest_palindromic_substring(str1 string) string {
	longest_palindromic_string := ""
	temp_str := ""

	runes := []rune(str1)

	low_index := 0
	high_index := low_index + 1

	//even palindrome.

	for low_index >= 0 && high_index <= len(str1)-1 {
		if runes[low_index] == runes[high_index] {
			temp_str = find_even_palindrome(runes, low_index, high_index, len(str1)-1)
			if len(temp_str) > len(longest_palindromic_string) {
				longest_palindromic_string = strings.Clone(temp_str)
			}
		}
		low_index += 1
		high_index += 1
	}

	low_index = 0
	high_index = low_index + 2
	index := low_index + 1

	//end_index
	for (low_index >= 0 && high_index <= (len(str1)-1)) && (index <= (len(str1) - 1)) {
		if runes[low_index] == runes[high_index] {
			temp_str = find_odd_palindrome(runes, low_index, high_index, len(str1))
			if len(temp_str) > len(longest_palindromic_string) {
				longest_palindromic_string = strings.Clone(temp_str)
			}
		}
		low_index += 1
		index = low_index + 1
		high_index = low_index + 2
	}

	return longest_palindromic_string
}

func main() {
	fmt.Println("Find the longest palindromic substring:")
	fmt.Println("Enter the string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 := scanner.Text()
	palndrome_str := find_longest_palindromic_substring(str1)
	fmt.Println(palndrome_str)
}
