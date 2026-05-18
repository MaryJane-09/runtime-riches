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

func TypeEffectFast(word []string) {
	for _, line := range word {
		for _, char := range line {
			fmt.Print(string(char))
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println()
		time.Sleep(900 * time.Millisecond)
	}
}

func TypeEffectSlow(word []string) {
	for _, line := range word {
		for _, char := range line {
			fmt.Print(string(char))
			time.Sleep(400 * time.Millisecond)
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