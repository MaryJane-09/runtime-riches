package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func play(input io.Reader, output io.Writer, questions []question) error {
	reader := bufio.NewScanner(input)
	fmt.Fprintln(output, "RUNTIME RICHES — LEEF CENTER EDITION SERIES 1")
	fmt.Fprintln(output, "Answer 20 Go questions to win ₦20,000,000 in-game.")
	fmt.Fprintln(output, "Enter HELP for rules, LIFELINE for help with a question, or WALKAWAY to keep your current prize.")

	var name string
	for name == "" {
		fmt.Fprint(output, "Your name: ")
		if !reader.Scan() {
			return reader.Err()
		}
		name = strings.TrimSpace(reader.Text())
	}

	for {
		if !playRound(reader, output, questions, name) {
			return reader.Err()
		}
		if !confirm(reader, output, "Play again?") {
			return reader.Err()
		}
	}
}

func playRound(reader *bufio.Scanner, output io.Writer, questions []question, name string) bool {
	currentPrize, safePrize := "0", "0"
	var used [3]bool

	for i, q := range questions {
		hidden := ""
		fmt.Fprintf(output, "\nQuestion %d of %d · Playing for ₦%s\n", i+1, len(questions), q.prize)
		fmt.Fprintf(output, "Current prize: ₦%s · Guaranteed: ₦%s\n", currentPrize, safePrize)
		showQuestion(output, q, hidden)

	answer:
		for {
			fmt.Fprint(output, "Answer: ")
			if !reader.Scan() {
				return false
			}
			answer := strings.ToUpper(strings.TrimSpace(reader.Text()))
			switch {
			case answer == "HELP":
				AskForHelp(output)
			case answer == "WALKAWAY":
				if confirm(reader, output, fmt.Sprintf("Walk away with ₦%s?", currentPrize)) {
					showResult(output, name, "You walked away.", currentPrize)
					return true
				}
			case answer == "LIFELINE":
				if useLifeline(reader, output, q, &used) {
					hidden = q.eliminated
					showQuestion(output, q, hidden)
				}
			case len(answer) != 1 || !strings.Contains(choices, answer):
				fmt.Fprintln(output, "Enter an option letter or type HELP.")
			case strings.Contains(hidden, answer):
				fmt.Fprintln(output, "That option was removed. Choose one of the remaining answers.")
			case answer[0] != q.answer:
				fmt.Fprintf(output, "Incorrect. The correct answer was %s.\n", q.options[q.answer-'A'])
				showResult(output, name, "Your round is over.", safePrize)
				return true
			default:
				currentPrize = q.prize
				fmt.Fprintf(output, "Correct! Your current prize is ₦%s.\n", currentPrize)
				switch i + 1 {
				case 5, 10, 15:
					safePrize = currentPrize
					fmt.Fprintf(output, "Safe point reached. ₦%s is now guaranteed.\n", safePrize)
				}
				break answer
			}
		}
	}

	showResult(output, name, "You answered every question correctly!", currentPrize)
	return true
}

func showQuestion(output io.Writer, q question, hidden string) {
	fmt.Fprintln(output, q.prompt)
	for i, option := range q.options {
		if !strings.ContainsRune(hidden, rune(choices[i])) {
			fmt.Fprintln(output, option)
		}
	}
}

func confirm(reader *bufio.Scanner, output io.Writer, prompt string) bool {
	for {
		fmt.Fprintf(output, "%s (yes/no): ", prompt)
		if !reader.Scan() {
			return false
		}
		switch strings.ToLower(strings.TrimSpace(reader.Text())) {
		case "yes", "y":
			return true
		case "no", "n":
			return false
		}
		fmt.Fprintln(output, "Enter yes or no.")
	}
}
