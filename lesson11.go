package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsack(arr, brr []int, n int, W int) int {
	if (n == 0) || (W == 0) {
		return 0
	} else if arr[n-1] > W {
		return knapsack(arr, brr, n-1, W)
	} else {
		return max(knapsack(arr, brr, n-1, W), knapsack(arr, brr, n-1, W-arr[n-1])+brr[n-1])
	}
}

func main() {
	fmt.Println("Knapsack problem:")
	arr := []int{1, 2, 3}
	brr := []int{4, 5, 1}
	fmt.Println(knapsack(brr, arr, 3, 4))
}
