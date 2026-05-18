package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Hints(i int) {

	data, err := os.ReadFile("hint.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := strings.Split(string(data), "\n")

	fmt.Println(slice[i-1])
}

func CreatorsSuggestions(i int) {
	data, err := os.ReadFile("suggestions.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := strings.Split(string(data), "\n")

	fmt.Println(slice[i-1])
}

func FiftyFifty(input []string, i int) []string {
	data, err := os.ReadFile("fiftyfifty.txt")
	if err != nil {
		fmt.Println(err)
	}
	slice := strings.Split(string(data), "\n")
	nums, _ := strconv.Atoi(string(slice[i-1][0]))
	num, _ := strconv.Atoi(string(slice[i-1][1]))
	input[nums] = ""
	input[num] = ""
	return input

}