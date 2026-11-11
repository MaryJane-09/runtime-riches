package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func useLifeline(reader *bufio.Scanner, output io.Writer, q question, used *[3]bool) bool {
	if used[0] && used[1] && used[2] {
		fmt.Fprintln(output, "You have used all three lifelines.")
		return false
	}

	for i, label := range []string{"50:50", "Hint", "Suggestion"} {
		status := ""
		if used[i] {
			status = " (used)"
		}
		fmt.Fprintf(output, "%d. %s%s\n", i+1, label, status)
	}
	fmt.Fprint(output, "Choose a lifeline or type BACK: ")
	if !reader.Scan() {
		return false
	}

	var selected int
	switch strings.ToUpper(strings.TrimSpace(reader.Text())) {
	case "1", "50:50":
		selected = 0
	case "2", "HINT":
		selected = 1
	case "3", "SUGGESTION":
		selected = 2
	case "BACK":
		return false
	default:
		fmt.Fprintln(output, "Unknown lifeline. Type LIFELINE to try again.")
		return false
	}

	if used[selected] {
		fmt.Fprintln(output, "You have already used that lifeline.")
		return false
	}
	used[selected] = true
	switch selected {
	case 0:
		fmt.Fprintln(output, "50:50 removed two wrong answers.")
		return true
	case 1:
		fmt.Fprintln(output, "Hint:", q.hint)
	case 2:
		fmt.Fprintln(output, "Suggestion:", q.suggestion)
	}
	return false
}
