package main

import "fmt"

func AskForHelp() {
	lines := []string{
		"",
		Color("GUARANTEED LEVELS:", "\033[1;35m"),
		Color("QUESTION 5  → SAFE POINT", "\033[1;35m"),
		Color("QUESTION 10 → HIGHER SAFE POINT", "\033[1;35m"),
		Color("QUESTION 15 → FINAL SAFE POINT", "\033[1;35m"),
		"",
		Color("LIFELINES AVAILABLE:", "\033[1;35m"),
		Color("1. 50:50 - REMOVE TWO WRONG OPTIONS", "\033[1;35m"),
		Color("2. HINT - SMALL CLUE", "\033[1;35m"),
		Color("3. SUGGESTION - GUIDED HELP", "\033[1;35m"),
		Color("ONCE USED... THEY CANNOT BE USED AGAIN.", "\033[1;35m"),
		"",
		Color("WALKAWAY - TO LEAVE WITH YOUR CURRENT PRIZE.", "\033[1;35m"),
		"",
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}