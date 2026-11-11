package main

import (
	"fmt"
	"io"
)

func AskForHelp(output io.Writer) {
	fmt.Fprintln(output, "Answer with A, B, C, or D. A wrong answer ends the round.")
	fmt.Fprintln(output, "Questions 5, 10, and 15 lock in guaranteed prizes once answered correctly.")
	fmt.Fprintln(output, "WALKAWAY lets you keep your current in-game prize after you confirm.")
	fmt.Fprintln(output, "Type LIFELINE to use 50:50, Hint, or Suggestion. Each is available once per round.")
	fmt.Fprintln(output, "50:50 removes two wrong answers. Treat the suggestion as advice.")
}
