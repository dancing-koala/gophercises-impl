package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dancing-koala/gophercises-impl/gophercise-1/pkg/questions"
	"io/ioutil"
	"os"
	"strings"
)

var (
	questionList []*questions.Question
	dataPath     string
)

const (
	DEFAULT_DATA_FILE = "./questions.csv"
)

func main() {
	flag.StringVar(&dataPath, "data-path", DEFAULT_DATA_FILE, "Path of the file containing the questions")
	flag.Parse()

	readQuestions(dataPath)

	right := 0
	input := bufio.NewReader(os.Stdin)

	printSeparator()

	for _, q := range questionList {
		attempt, _ := input.ReadString('\n')

		printQuestion(q)

		if q.VerifyAnswer(strings.Trim(attempt, "\n")) {
			right++
		}
	}

	printSeparator()
	printScore(right, len(questionList))
	printSeparator()
}

func printSeparator() {
	fmt.Println("---")
}

func printQuestion(q *questions.Question) {
	fmt.Printf("%s=", q.Text)
}

func printScore(right, total int) {
	fmt.Printf("You scored %d out of %d\n", right, total)
}

func readQuestions(csvPath string) {
	rawData, err := ioutil.ReadFile(csvPath)

	if err != nil {
		panic(err)
	}

	questionList = questions.ReadCsv(rawData)
}
