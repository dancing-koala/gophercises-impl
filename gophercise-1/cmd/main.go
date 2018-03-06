package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type question struct {
	text   string
	answer string
}

var (
	questions []*question
)

func main() {
	readQuestions()

	total := len(questions)
	good := 0

	input := bufio.NewReader(os.Stdin)

	printSeparator()

	for _, q := range questions {
		fmt.Printf("%s=", q.text)
		attempt, _ := input.ReadString('\n')

		if checkAnswer(q, strings.Trim(attempt, "\n")) {
			good++
		}
	}

	printSeparator()
	fmt.Printf("You scored %d out of %d\n", good, total)
	printSeparator()
}

func printSeparator() {
	fmt.Println("---")
}

func checkAnswer(q *question, attempt string) bool {
	return strings.Compare(q.answer, attempt) == 0
}

func readQuestions() {
	rawData, err := ioutil.ReadFile("./questions.csv")

	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(bytes.NewReader(rawData))

	for {
		row, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		questions = append(questions, &question{
			text:   row[0],
			answer: row[1],
		})
	}
}
