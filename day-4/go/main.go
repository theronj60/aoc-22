package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func cleanString(str string) string {
	var regX = regexp.MustCompile(`[^a-zA-Z0-9 - ]+`)
	return regX.ReplaceAllString(str, " ")
}

func getNew(split []string) [][]string {
	var arr [][]string
	var tempArray []string
	for _, value := range split {
		tempArray = strings.Fields(cleanString(value))
		arr = append(arr, tempArray)
	}
	return arr
}

func main() {
	// content, err := ioutil.ReadFile("../section.test")
	content, err := ioutil.ReadFile("../section.prod")
	if err != nil {
		fmt.Println(err, "Could not find file.")
	}

	split := strings.Fields(string(content))

	newArray := getNew(split)
	
	var count int

	for _, item := range newArray {
		x1, _ := strconv.Atoi(string(item[0]))
		x2, _ := strconv.Atoi(string(item[1]))
		y1, _ := strconv.Atoi(string(item[2]))
		y2, _ := strconv.Atoi(string(item[3]))

		if x1 <= y1 && x2 >= y2 || y1 <= x1 && y2 >= x2 {
			count++
		}
	}
	fmt.Println("part one: ", count)
}
