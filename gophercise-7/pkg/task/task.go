package task

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"strconv"
)

type Task struct {
	Id   uint64
	Name string
	Done bool
}

var bucketName = []byte("TODO")

type TaskRepository struct {
	db *bolt.DB
}

func itob(val uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	return b
}

func NewRepository(db *bolt.DB) *TaskRepository {
	if db == nil {
		return nil
	}

	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) ListNotDoneTasks() ([]Task, error) {
	return listTasks(tr.db, func(task Task) bool {
		return !task.Done
	})
}

func listTasks(db *bolt.DB, filter func(Task) bool) ([]Task, error) {
	tasks := make([]Task, 0, 32)

	dbErr := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		if b == nil {
			return nil
		}

		c := b.Cursor()

		for key, val := c.First(); key != nil; key, val = c.Next() {

			task := &Task{}

			json.Unmarshal(val, &task)

			if task.Done {
				continue
			}

			tasks = append(tasks, *task)
		}

		return nil
	})

	return tasks, dbErr
}

func (tr *TaskRepository) AddTask(taskName string) error {
	dbErr := tr.db.Update(func(tx *bolt.Tx) error {

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

	return dbErr
}

func (tr *TaskRepository) DoTask(taskId string) (*Task, error) {

	task := &Task{}

	dbErr := tr.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(bucketName)

		if b == nil {
			return errors.New("Task not found: " + taskId)
		}

		id, err := strconv.ParseUint(taskId, 10, 64)

		if err != nil {
			return err
		}

		key := itob(id)
		val := b.Get(key)

		if val == nil {
			return errors.New("Task not found: " + taskId)
		}

		json.Unmarshal(val, &task)

		if task.Done {
			return errors.New(fmt.Sprintf("Task %q already done!\n", task.Name))
		}

		task.Done = true

		buf, err := json.Marshal(task)

		if err != nil {
			return err
		}

		b.Put(key, buf)

		return nil
	})

	return task, dbErr
}
