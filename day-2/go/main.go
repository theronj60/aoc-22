package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// opponent
// A => rock
// B => paper
// C => scissor

// player
// X => rock
// Y => paper
// Z => scissor

func getHand(hand string) string {
	var gesture string
	switch hand {
	case "A", "X":
		gesture = "rock"
	case "B", "Y":
		gesture = "paper"
	case "C", "Z":
		gesture = "scissor"
	}
	return gesture
}

func main() {
	content, err := ioutil.ReadFile("../rps.test")
	// content, err := ioutil.ReadFile("../rps.prod")
	if err != nil {
		fmt.Println("Err")
	}

	var total int

	// scoring based on shape
	// rock = 1
	// paper = 2
	// scissors = 3
	// scoring based on outcome
	// loss = 0
	// draw = 3
	// win = 6

	split := strings.Split(string(content), "\n")
	for _, item := range split {
		// second item in slice is a space
		// compare item[0] (opponent) to item[2] (player)
		fmt.Println(string(item[1]))
	}
}
