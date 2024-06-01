package main

import "fmt"

func is_shuffle(str1, str2, str3 string) bool {
	lngth := len(str1) + len(str2)

	fmt.Println(lngth)

	if lngth != len(str3) {
		return false
	}

	index1 := 0
	index2 := 0
	index := 0

	for (index1 < len(str1)) && (index2 < len(str2)) {
		if str1[index1] == str3[index] {
			index1 += 1
			index += 1
		} else if str2[index2] == str3[index] {
			index2 += 1
			index += 1
		} else {
			return false
		}
	}

	fmt.Println(index, index1, index2)

	if index1 == len(str1) {
		for index2 < len(str2) {
			if str3[index] == str2[index2] {
				index += 1
				index2 += 1
			} else {
				return false
			}
		}
	} else if index2 == len(str2) {
		for index1 < len(str1) {
			if str1[index1] == str3[index] {
				index += 1
				index1 += 1
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println("Find if the string is the  shuffle of the given string")
	str1 := "xy"
	str2 := "12"
	str3 := "x1y2"
	test := is_shuffle(str1, str2, str3)
	fmt.Println(test)
}
