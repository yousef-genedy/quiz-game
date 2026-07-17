package main

import (
	"encoding/csv"
	"os"
)

func ReadProblems(f string) [][]string {
	file, err := os.Open(f)
	if err != nil {
		Exit(err.Error())
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			Exit(err.Error())
		}
	}(file)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		Exit(err.Error())
	}

	return data
}
