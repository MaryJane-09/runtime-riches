package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestFullRoundAndReplay(t *testing.T) {
	questions, err := loadQuestions()
	if err != nil {
		t.Fatal(err)
	}
	var input strings.Builder
	input.WriteString("MJ\n")
	for _, q := range questions {
		fmt.Fprintf(&input, "%c\n", q.answer)
	}
	input.WriteString("yes\nwalkaway\nyes\nno\n")

	var output bytes.Buffer
	if err := play(strings.NewReader(input.String()), &output, questions); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"Question 20 of 20",
		"In-game winnings: ₦20,000,000",
		"In-game winnings: ₦0",
	} {
		if !strings.Contains(output.String(), want) {
			t.Errorf("missing %q in game output", want)
		}
	}
}

func TestWrongAnswerKeepsGuaranteedPrize(t *testing.T) {
	questions, err := loadQuestions()
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		correct int
		prize   string
	}{
		{0, "0"},
		{5, "20,000"},
		{10, "500,000"},
		{15, "3,000,000"},
	} {
		t.Run(tc.prize, func(t *testing.T) {
			var input strings.Builder
			input.WriteString("Mary Jane\n")
			for _, q := range questions[:tc.correct] {
				fmt.Fprintf(&input, "%c\n", q.answer)
			}
			fmt.Fprintf(&input, "%c\nno\n", questions[tc.correct].eliminated[0])

			var output bytes.Buffer
			if err := play(strings.NewReader(input.String()), &output, questions); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(output.String(), "In-game winnings: ₦"+tc.prize+"\n") {
				t.Errorf("expected guaranteed prize ₦%s", tc.prize)
			}
		})
	}
}

func TestWalkawayAndEliminatedAnswers(t *testing.T) {
	questions, err := loadQuestions()
	if err != nil {
		t.Fatal(err)
	}
	first := questions[0]
	input := fmt.Sprintf("MJ\nwalkaway\nmaybe\nno\nlifeline\n1\n%c\nlifeline\n50:50\n%c\nwalkaway\nyes\nno\n",
		first.eliminated[0], first.answer)

	var output bytes.Buffer
	if err := play(strings.NewReader(input), &output, questions); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"Enter yes or no.",
		"That option was removed.",
		"You have already used that lifeline.",
		"In-game winnings: ₦1,000",
	} {
		if !strings.Contains(output.String(), want) {
			t.Errorf("missing %q in game output", want)
		}
	}
}
