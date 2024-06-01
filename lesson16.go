package main

import "fmt"

func find_min_cost(word []int, rem int, index int, K int) int {
	if index == len(word) {
		return 0
	}
	if word[index] > rem {
		return (rem+1)*(rem+1) + find_min_cost(word, K-word[index]-1, index+1, K)
	} else {
		return min((rem+1)*(rem+1)+find_min_cost(word, K-word[index]-1, index+1, K), find_min_cost(word, rem-word[index]-1, index+1, K))
	}
}

func main() {
	fmt.Println("Find the minimum cost for wrapping word:")
	word := []int{3, 2, 2}
	var K int
	K = 4
	cost := find_min_cost(word, K, 0, K)
	fmt.Println("The total cost of word wrapping:", cost)
}
