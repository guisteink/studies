package structures

import (
	"errors"
	"fmt"
)

type Queue struct {
	data []byte
}

func (q *Queue) Enqueue(data []byte) error {
	fmt.Println("Inserindo: ", data)

	for _, b := range data {
		q.data = append(q.data, b)
	}

	fmt.Println("Queue: ", q.data)

	return nil
}

func (q *Queue) Dequeue() (byte, error) {
	if len(q.data) == 0 {
		return 0, errors.New("queue is empty")
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) GetSize() int {
	return len(q.data)
}

func (q *Queue) Peek() (byte, error) {
	if len(q.data) == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue) Clear() error {
	q.data = nil
	return nil
}
