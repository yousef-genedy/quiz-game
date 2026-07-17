package main

import "flag"

func ReadFlags() (*string, *int, *bool) {
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the quiz order")

	flag.Parse()

	return csvFile, limit, shuffle
}
