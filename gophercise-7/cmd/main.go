package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"os"
	"strconv"
	"strings"
)

var (
	db         *bolt.DB
	bucketName = []byte("TODO")
)

type Task struct {
	Id   uint64
	Name string
	Done bool
}

func main() {
	if len(os.Args) < 2 {
		printMan()
		return
	}

	var err error

	db, err = bolt.Open("./db/gophercise_7.db", 0600, nil)
	defer db.Close()

	handleErr(err)

	command := os.Args[1]

	switch command {
	case "list":
		listTasks()
		break

	case "add":
		taskName := strings.Join(os.Args[2:], " ")
		addTask(taskName)
		break

	case "do":
		taskId := os.Args[2]
		doTask(taskId)
		break

	default:
		handleErr(errors.New("Unknow command: " + command))
		return
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
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

func itob(val uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	return b
}

func listTasks() {
	dbErr := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		if b == nil {
			fmt.Println("No task in your list.")
			return errors.New("Bucket does not exist!")
		}

		fmt.Println("You have the following tasks:")

		c := b.Cursor()

		task := &Task{}

		for key, val := c.First(); key != nil; key, val = c.Next() {

			json.Unmarshal(val, &task)

			if task.Done {
				continue
			}

			fmt.Printf("%d. %s\n", task.Id, task.Name)
		}

		return nil
	})

	handleErr(dbErr)
}

func addTask(taskName string) {
	dbErr := db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(bucketName)

		var err error

		if b == nil {
			b, err = tx.CreateBucket(bucketName)

			if err != nil {
				return err
			}
		}

		id, _ := b.NextSequence()

		task := &Task{
			Id:   id,
			Name: taskName,
			Done: false,
		}

		buf, err := json.Marshal(task)

		if err != nil {
			return err
		}

		b.Put(itob(id), buf)

		return nil
	})

	handleErr(dbErr)

	fmt.Printf("Added %q to your task list.\n", taskName)
}

func doTask(taskId string) {
	dbErr := db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(bucketName)

		if b == nil {
			fmt.Println("Task not found")
			return errors.New("Bucket does not exist!")
		}

		id, err := strconv.ParseUint(taskId, 10, 64)

		if err != nil {
			return err
		}

		key := itob(id)
		val := b.Get(key)

		if val == nil {
			fmt.Printf("Task <%s> does not exists!\n", taskId)
			return errors.New("Task not found: " + taskId)
		}

		task := &Task{}

		json.Unmarshal(val, &task)

		if task.Done {
			fmt.Printf("Task %q already done!\n", task.Name)
			return errors.New("Task already done")
		}

		task.Done = true

		buf, err := json.Marshal(task)

		if err != nil {
			return err
		}

		b.Put(key, buf)

		fmt.Printf("You have completed the %q task.\n", task.Name)

		return nil
	})

	handleErr(dbErr)
}
