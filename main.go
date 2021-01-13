package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// The problem - (question/answer)
type problem struct {
	q string
	a string
}

func main() {
	// Sets command line flags
	csvFilename := flag.String("csv", "problems.csv",
		"a csv file in the format 'question,answer'")
	flag.Parse()

	// Opens the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
	}

	// Creates an instance of the csv reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file")
	}

	fmt.Println(lines)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

// Helper function "exit"
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
