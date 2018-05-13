package main

import (
	"flag"
	"fmt"
	"github.com/dancing-koala/gophercises-impl/gophercise-1/pkg/questions"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type questionList []*questions.Question

const (
	DEFAULT_DATA_FILE      = "./problems.csv"
	DEFAULT_TIMER_DURATION = "30s"
	DEFAULT_SHUFFLE        = false
)

var (
	list          questionList
	dataPath      string
	timerDuration string
	shuffle       bool
)

func main() {
	flag.StringVar(&dataPath, "data-path", DEFAULT_DATA_FILE, "Path of the file containing the questions")
	flag.StringVar(&timerDuration, "timer-duration", DEFAULT_TIMER_DURATION, "Duration of the timer, you can use h,m,s,ms,us & ns symbols")
	flag.BoolVar(&shuffle, "shuffle", DEFAULT_SHUFFLE, "Indicates whether the question list should be shuffled or not")
	flag.Parse()

	readQuestions(dataPath)

	if shuffle {
		shuffleQuestions()
	}

	printSeparator()

	d, err := time.ParseDuration(timerDuration)

	if err != nil {
		panic(err)
	}

	timer := time.NewTimer(d)

	right := 0

	for _, q := range list {

		answerChan := make(chan string, 1)
		printQuestion(q)

		go func() {
			var a string
			fmt.Scanln(&a, '\n')
			answerChan <- a
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTIMEOUT !!")
			printScore(right, len(list))
			return

		case attempt := <-answerChan:
			if q.VerifyAnswer(strings.Trim(attempt, "\n")) {
				right++
			}

			break
		}
	}

	timer.Stop()

	printScore(right, len(list))
}

func printSeparator() {
	fmt.Println("---")
}

func printQuestion(q *questions.Question) {
	fmt.Printf("%s = ", q.Text)
}

func printScore(right, total int) {
	printSeparator()
	fmt.Printf("You scored %d out of %d\n", right, total)
	printSeparator()
}

func readQuestions(csvPath string) {
	data, err := ioutil.ReadFile(csvPath)

	if err != nil {
		panic(err)
	}

	list = questions.ReadCsv(data)
}

func shuffleQuestions() {
	size := len(list)
	rand.Seed(time.Now().Unix())

	for i := 0; i < size; i++ {
		dest := rand.Intn(size)

		if dest != i {
			tmp := list[dest]
			list[dest] = list[i]
			list[i] = tmp
		}
	}
}
