package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=====================")
	str1 := "hello world"
	list_of_str := strings.Split(str1, " ")
	fmt.Printf("list of string %v and type is %T \n", list_of_str, list_of_str)
	for index, value := range list_of_str {
		fmt.Println(index, value)
	}

}
