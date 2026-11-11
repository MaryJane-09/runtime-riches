package main

import (
	"fmt"
	"os"
)

func main() {
	questions, err := loadQuestions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot load Runtime Riches:", err)
		os.Exit(1)
	}

	if err := play(os.Stdin, os.Stdout, questions); err != nil {
		fmt.Fprintln(os.Stderr, "Runtime Riches:", err)
		os.Exit(1)
	}
}
