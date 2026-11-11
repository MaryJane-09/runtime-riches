package main

import (
	"fmt"
	"io"
)

func showResult(output io.Writer, name, outcome, prize string) {
	fmt.Fprintf(output, "\n%s, %s\n", name, outcome)
	fmt.Fprintf(output, "In-game winnings: ₦%s\n", prize)
}
