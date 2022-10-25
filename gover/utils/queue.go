package utils

import "errors"

// only for 1 read and 1 write model
type Queue[T any] struct {
	c   chan T
	cap int
}

func (q *Queue[T]) Enqueue(data T) error {
	if len(q.c) >= q.cap {
		return errors.New("queue is full")
	}
	q.c <- data
	return nil
}

func (q *Queue[T]) Dequeue() T {
	return <-q.c
}

func (q *Queue[T]) Close() { close(q.c) }

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		c:   make(chan T, size),
		cap: size,
	}
}
