package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var alphabet = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func compare(word string, value string) bool {
	for i, letter := range word {
		if i > len(word)/2 - 1 {
			if string(letter) == value {
				return true
			}
		}
	}
	return false
}

func getInt(value string) int {
	var num int
	for i, item := range alphabet {
		if item == value {
			num = i + 1
			return num
		}
	}
	return num
}

func main() {
	content, err := ioutil.ReadFile("../ruck.test")
	// content, err := ioutil.ReadFile("../ruck.prod")
	if err != nil {
		fmt.Println("Err")
	}

	split := strings.Split(string(content), "\n")
	
	var sum int
	
	for _, value := range split {
		var letter string
		for i, item := range value {
			if i < len(value)/2 {
				if compare(value, string(item)) {
					letter = string(item)
				}
			}
		}
		fmt.Println(letter)
		sum += getInt(letter)
	}
	fmt.Println("sum part one: ", sum)
}
