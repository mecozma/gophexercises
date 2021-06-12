package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// declare the csvFile flag.
	csvFile := flag.String("csv", "challenges.scv", `a csv file in format 
		of 'question,answer`)
	flag.Parse()

	// open the csv file and check for errors.
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFile))
	}

	// instantiate a new reader that reads the CSV file.
	r := csv.NewReader(file)

	// read all lines of the csv file.
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}

	// instantiate a slice of challenges.
	challenges := parseLines(lines)

	// print the quiz answers.
	printQuestion(challenges)

}
