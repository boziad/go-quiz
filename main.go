package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Problem struct {
	q string
	a string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")

	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		log.Fatalf("failed to open the file %s\n", *csvFileName)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		log.Fatalf("failed to parse the file %s\n", *csvFileName)
	}

	probelms := parseLines(lines)
	correct := 0
	for i, p := range probelms {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(probelms))

}

func parseLines(lines [][]string) (ret []Problem) {
	ret = make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{q: line[0], a: strings.TrimSpace(line[1])}
	}

	return

}
