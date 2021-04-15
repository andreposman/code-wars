/*
In this Kata, you will be given a string and your task will be to return
a list of ints detailing the count of uppercase letters, lowercase, numbers and special characters, as follows.

Solve("*'&ABCDabcde12345") = [4,5,5,3].
 --the order is: uppercase letters, lowercase, numbers and special characters.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {

	s := "*'&ABCDabcde12345"

	Solve(s)
}

func Solve(s string) []int {
	var upper []string
	var lower []string
	var numbers []string
	var special []string
	var countSlice []int

	for i := 0; i < len(s); i++ {
		str := string(s)
		if isUpper(str[i]) && !isSpecial(str[i]) && !isNumber(str[i]) {
			upper = append(upper, string(str[i]))
		}
		if !isUpper(str[i]) && !isSpecial(str[i]) && !isNumber(str[i]) {
			lower = append(lower, string(str[i]))
		}
		if isNumber(str[i]) {
			numbers = append(numbers, string(str[i]))
		}
		if isSpecial(str[i]) {
			special = append(special, string(str[i]))
		}
	}

	countSlice = append(countSlice, len(upper), len(lower), len(numbers), len(special))
	fmt.Println(countSlice)

	return countSlice
}

func isUpper(s byte) bool {
	for _, v := range string(s) {
		if !unicode.IsUpper(v) && unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func isNumber(s byte) bool {
	for _, v := range string(s) {
		if unicode.IsNumber(v) {
			return true
		}
	}
	return false
}

func isSpecial(s byte) bool {
	for _, v := range string(s) {
		if !unicode.IsUpper(v) && !unicode.IsLower(v) && !unicode.IsNumber(v) {
			return true
		}
	}
	return false
}
