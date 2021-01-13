package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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

	problems := parseLines(lines)
	printAndCheck(problems)
}

func printAndCheck(problems []problem) {
	// Print questions to stdout and check answer
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// Helper function "exit"
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
