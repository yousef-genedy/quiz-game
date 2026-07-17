package main

import (
	"fmt"
	"strings"
)

func Process(problems [][]string, score *Score) {
	for i, problem := range problems {
		var input string
		question, answer := problem[0], strings.TrimSpace(problem[1])

		fmt.Printf("Problem #%d: %s = ", i+1, question)

		_, err := fmt.Scan(&input)
		if err != nil {
			Exit(err.Error())
		}

		score.Add(input == answer)
	}
}
