package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Color(input, text string) string {
	return text + input + "\033[0m"
}

func ClearScreen() {
	command := exec.Command("clear")
	command.Stdout = os.Stdout
	command.Run()
}

func TypeEffectFast(words []string) {
	for _, line := range words {
		for _, char := range line {
			fmt.Print(string(char))

			if char != ' ' {
				time.Sleep(150 * time.Millisecond)
			}
		}
		fmt.Println()
		time.Sleep(90 * time.Millisecond)
	}
}

func TypeEffectSlow(words []string) {
	for _, line := range words {
		for _, char := range line {
			fmt.Print(string(char))

			if char != ' ' {
				time.Sleep(150 * time.Millisecond)
			}
		}
		fmt.Println()
		time.Sleep(900 * time.Millisecond)
	}
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		word[i] = strings.ToUpper(string(word[i][0])) + strings.ToLower(word[i][1:])
	}
	return strings.Join(word, " ")
}

