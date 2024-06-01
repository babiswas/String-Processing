package main

import "fmt"

func isShuffle(str1, str2, str3 string) bool {
	l1 := len(str1)
	l2 := len(str2)
	l3 := len(str3)
	index1 := 0
	index2 := 0
	k := 0
	for k < l3 {
		if index1 < l1 && str1[index1] == str3[k] {
			index1 += 1
		} else if index2 < l2 && str2[index2] == str3[k] {
			index2 += 1
		} else {
			break
		}
		k += 1
	}
	if index1 < l1 || index2 < l2 {
		return false
	}
	return true
}

func main() {
	fmt.Println("Verify if the string is a shuffle of another 2 strings:")
	if isShuffle("xy", "12", "x1y2") {
		fmt.Println("The string is a shuffle of the other 2 strings.")
	} else {
		fmt.Println("The string is not a shuffle of the other 2 strings.")
	}
}
