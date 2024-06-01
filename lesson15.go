package main

import (
	"bufio"
	"fmt"
	"os"
)

func permute(rns []rune, mp map[string]bool) []string {
	glob_str := make([]string, 0)
	if len(rns) == 1 {
		test := make([]string, 0)
		test = append(test, string(rns[len(rns)-1]))
		return test
	}
	for i := 0; i < len(rns); i++ {
		sample := rns[len(rns)-1]
		rns = rns[:len(rns)-1]
		perms := permute(rns, mp)
		for index, _ := range perms {
			perms[index] += string(sample)
		}
		for _, val := range perms {
			glob_str = append(glob_str, val)
		}
		rns = append([]rune{sample}, rns...)

	}
	return glob_str
}

func main() {
	fmt.Println("Find permutation of a string:")
	fmt.Println("Enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 := scanner.Text()
	mp := make(map[string]bool)
	all_permutation := permute([]rune(str1), mp)
	fmt.Println(all_permutation)
}
