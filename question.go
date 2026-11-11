package main

import (
	_ "embed"
	"fmt"
	"strings"
)

const questionCount = 20
const choices = "ABCD"

//go:embed question.txt
var questionData string

//go:embed hint.txt
var hintData string

//go:embed suggestions.txt
var suggestionData string

//go:embed fifty_fifty.txt
var fiftyFiftyData string

type question struct {
	prompt     string
	options    [4]string
	answer     byte
	prize      string
	hint       string
	suggestion string
	eliminated string
}

func dataLines(data string) []string {
	return strings.Split(strings.TrimRight(strings.ReplaceAll(data, "\r\n", "\n"), "\n"), "\n")
}

func loadQuestions() ([]question, error) {
	lines := dataLines(questionData)
	hints := dataLines(hintData)
	suggestions := dataLines(suggestionData)
	pairs := dataLines(fiftyFiftyData)
	if len(lines) != questionCount*8-1 || len(hints) != questionCount ||
		len(suggestions) != questionCount || len(pairs) != questionCount {
		return nil, fmt.Errorf("expected %d questions with matching lifelines", questionCount)
	}

	questions := make([]question, 0, questionCount)
	for i := 0; i < questionCount; i++ {
		start := i * 8
		fields := lines[start : start+7]
		if i < questionCount-1 && strings.TrimSpace(lines[start+7]) != "" {
			return nil, fmt.Errorf("missing separator after question %d", i+1)
		}
		for _, field := range fields {
			if strings.TrimSpace(field) == "" {
				return nil, fmt.Errorf("empty field in question %d", i+1)
			}
		}
		answer := fields[5]
		if len(answer) != 1 || !strings.Contains(choices, answer) {
			return nil, fmt.Errorf("invalid answer for question %d", i+1)
		}

		pair := strings.TrimSpace(pairs[i])
		if len(pair) != 2 || pair[0] < '1' || pair[0] > '4' ||
			pair[1] < '1' || pair[1] > '4' || pair[0] == pair[1] {
			return nil, fmt.Errorf("invalid 50:50 choices for question %d", i+1)
		}
		eliminated := string([]byte{choices[pair[0]-'1'], choices[pair[1]-'1']})
		if strings.Contains(eliminated, answer) {
			return nil, fmt.Errorf("50:50 removes the correct answer for question %d", i+1)
		}
		if strings.TrimSpace(hints[i]) == "" || strings.TrimSpace(suggestions[i]) == "" {
			return nil, fmt.Errorf("missing lifeline text for question %d", i+1)
		}
		questions = append(questions, question{
			prompt: fields[0], options: [4]string(fields[1:5]), answer: answer[0],
			prize: fields[6], hint: hints[i], suggestion: suggestions[i], eliminated: eliminated,
		})
	}
	return questions, nil
}
