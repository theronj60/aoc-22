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
		// below if block, addresses an empty last item
		if item == split[len(split)-1] && len(split) > 100 {
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
	fmt.Println("Part One: ", total)
	partTwo(split)
}

func getResult(item string) string {
	var result string
	switch item {
	case "X":
		result = "lose"
	case "Y":
		result = "draw"
	case "Z":
		result = "win"
	}
	return result
}

func partTwo(list []string) {
	var opponent string
	var result string
	var score int
	var total int

	for _, item := range list {
		// below if block, addresses an empty last item
		if item == list[len(list)-1] && len(list) > 100 {
			break
		}

		opponent = getHand(string(item[0]))
		result = getResult(string(item[2]))

		rock := 1
		paper := 2
		scissor := 3

		if opponent == "rock" {
			if result == "draw" {
				score = 3 + rock
			} else if result == "win" {
				score = 6 + paper
			} else {
				score = scissor
			}
		} else if opponent == "paper" {
			if result == "draw" {
				score = 3 + paper
			} else if result == "win" {
				score = 6 + scissor
			} else {
				score = rock
			}
		} else if opponent == "scissor" {
			if result == "draw" {
				score = 3 + scissor
			} else if result == "win" {
				score = 6 + rock
			} else {
				score = paper
			}
		}
		total += score
	}
	fmt.Println("Part Two: ", total)
}
