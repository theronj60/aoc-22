package main

import (
	"fmt"
	"io/ioutil"
)

func compare(str string) bool {
	// pass array of four
	for i, letter := range str {
		for iter, item := range str {
			if i != iter {
				if item == letter {
					return false
				}
				continue
			}
		}
	}
	return true
}

func main() {
	content, err := ioutil.ReadFile("../buffer.prod")
	// content, err := ioutil.ReadFile("../buffer.test")
	if err != nil {
		fmt.Println(err, "Could not find file.")
	}

	split := string(content)
	var answer int
	var answerTwo int

	for i := range split {
		if i+1 < len(split)-2 {
			endFour := i + 4
			if compare(split[i:endFour]) {
				answer = endFour
				break
			}
			
		}
		
	}
	for i := range split {
		if i+1 < len(split)-20 {
			endBig := i + 14
			if compare(split[i:endBig]) {
				answerTwo = endBig
				break
			}
		}
	}
	fmt.Println("Part One: ", answer)
	fmt.Println("Part Two: ", answerTwo)
}
