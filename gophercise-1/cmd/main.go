package main

import (
	"bufio"
	"fmt"
	"github.com/dancing-koala/gophercises-impl/gophercise-1/pkg/questions"
	"io/ioutil"
	"os"
	"strings"
)

var (
	questionList []*questions.Question
)

func main() {
	readQuestions("./questions.csv")

	total := len(questionList)
	good := 0

	input := bufio.NewReader(os.Stdin)

	printSeparator()

	for _, q := range questionList {
		fmt.Printf("%s=", q.Text)
		attempt, _ := input.ReadString('\n')

		if q.VerifyAnswer(strings.Trim(attempt, "\n")) {
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

func readQuestions(csvPath string) {
	rawData, err := ioutil.ReadFile(csvPath)

	if err != nil {
		panic(err)
	}

	questionList = questions.ReadCsv(rawData)
}
