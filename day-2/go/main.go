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

// var total int
// total = 6
// scoring based on shape
// rock = 1
// paper = 2
// scissors = 3
// scoring based on outcome
// loss = 0
// draw = 3
// win = 6

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
	// content, err := ioutil.ReadFile("../rps.test")
	content, err := ioutil.ReadFile("../rps.prod")
	if err != nil {
		fmt.Println("Err")
	}

	split := strings.Split(string(content), "\n")

	var opponent string
	var player string
	var score int
	var total int

	for _, item := range split {
		// below if addresses an empty last item
		if item == split[len(split)-1] {
			break
		}
		opponent = getHand(string(item[0]))
		player = getHand(string(item[2]))
		switch player {
		case "rock":
			score = 1
		case "paper":
			score = 2
		case "scissor":
			score = 3
		}
		if opponent == "rock" {
			if player == "rock" {
				score += 3
			} else if player == "paper" {
				score += 6
			}
		} else if opponent == "paper" {
			if player == "paper" {
				score += 3
			} else if player == "scissor" {
				score += 6
			}
		} else if opponent == "scissor" {
			if player == "scissor" {
				score += 3
			} else if player == "rock" {
				score += 6
			}
		}
		total += score
	}
	fmt.Println(total)
}
