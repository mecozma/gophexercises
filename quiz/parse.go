package main

import (
	"fmt"
	"os"
	"strings"
)

// struct with the question and answer.
type problem struct {
	question string
	answer   string
}

// parse the csv file's lines.
func parseLines(lines [][]string) []problem {
	data := make([]problem, len(lines))
	for i, line := range lines {
		data[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return data
}

// exit function implements os.Exit().
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// print the quiz questions and check if the answer is right.
func printQuestion(challenges []problem) {
	countAnswers := 0
	for i, c := range challenges {
		fmt.Printf("Question #%d : %s = \n", i+1, c.question)

		// get the user's answer.
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == c.answer {
			countAnswers++
		}

	}
	quizScore(countAnswers, challenges)
}

// print quiz's score.
func quizScore(score int, challenges []problem) {
	fmt.Printf("You scored %d from %d questions!\n", score, len(challenges))
}
