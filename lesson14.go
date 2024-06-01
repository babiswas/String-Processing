package main

import (
	"bufio"
	"fmt"
	"os"
)

func longest_repeating_subsequence(rns1 []rune, rns2 []rune, n int, m int) int {
	if m-1 < 0 || n-1 < 0 {
		return 0
	}
	if (rns1[n-1] == rns2[m-1]) && (m-1 != n-1) {
		return 1 + longest_repeating_subsequence(rns1, rns2, n-1, m-1)
	} else {
		return max(longest_repeating_subsequence(rns1, rns2, n, m-1), longest_repeating_subsequence(rns1, rns2, n-1, m))
	}
}

func main() {
	fmt.Println("Find the longest repeating subsequence.")

	fmt.Println("Enter the string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 := scanner.Text()

	runes := []rune(str1)

	longest_subsequence := longest_repeating_subsequence(runes, runes, len(str1), len(str1))
	fmt.Println("Length of the longest repeating subsequence is:", longest_subsequence)
}
