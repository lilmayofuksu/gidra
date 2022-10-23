package utils

import "errors"

// only for 1 read and 1 write model
type BytesQueue struct {
	c   chan []byte
	cap int
}

func (q *BytesQueue) Enqueue(data []byte) error {
	if len(q.c) >= q.cap {
		return errors.New("queue is full")
	}
	q.c <- data
	return nil
}

func (q *BytesQueue) Dequeue() (ret []byte) {
	return <-q.c
}

func (q *BytesQueue) Close() { close(q.c) }

func NewBytesQueue(size int) *BytesQueue {
	return &BytesQueue{
		c:   make(chan []byte, size),
		cap: size,
	}
}
