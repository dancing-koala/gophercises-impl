package main

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/dancing-koala/gophercises-impl/gophercise-7/pkg/task"
	"os"
	"strings"
)

var repo *task.TaskRepository

func main() {
	if len(os.Args) < 2 {
		printMan()
		return
	}

	db, err := bolt.Open("./db/gophercise_7.db", 0600, nil)
	defer db.Close()

	repo = task.NewRepository(db)

	handleErr(err)

	command := os.Args[1]

	switch command {
	case "list":
		listNotDoneTasks()
		break

	case "completed":
		listDoneTasks()
		break

	case "add":
		taskName := strings.Join(os.Args[2:], " ")
		addTask(taskName)
		break

	case "do":
		taskId := os.Args[2]
		doTask(taskId)
		break

	case "rm":
		taskId := os.Args[2]
		removeTask(taskId)
		break

		break

	default:
		handleErr(errors.New("Unknow command: " + command))
		return
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}

func printMan() {
	fmt.Println(`task is a CLI for managing your TODOs.

Usage:
	task [command]

Available Commands:
	add     Add a new task to your TODO list
	do      Mark a task on your TODO list as complete
	list    List all of your incomplete tasks

Use "task [command] --help" for more information about a command.
`)
}

func listNotDoneTasks() {
	tasks, err := repo.ListNotDoneTasks()

	handleErr(err)

	if len(tasks) == 0 {
		fmt.Println("You have no task in your list.")
		return
	}

	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.Id, task.Name)
	}
}

func listDoneTasks() {
	tasks, err := repo.ListDoneTasks()

	handleErr(err)

	if len(tasks) == 0 {
		fmt.Println("You have no completed task in your list.")
		return
	}

	fmt.Println("You have the following tasks:")

	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.Id, task.Name)
	}
}

func addTask(taskName string) {
	err := repo.AddTask(taskName)

	handleErr(err)

	fmt.Printf("Added %q to your task list.\n", taskName)
}

func removeTask(taskId string) {
	task, err := repo.RemoveTask(taskId)

	handleErr(err)

	fmt.Printf("You have removed task %q.\n", task.Name)
}

func doTask(taskId string) {
	task, err := repo.DoTask(taskId)

	handleErr(err)

	fmt.Printf("You have completed the %q task.\n", task.Name)
}
