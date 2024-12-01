package queue

import "fmt"

type Queue[T any] struct {
	items []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) DeQueue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue[T]) Size() int {
	return (len(q.items))
}

func (q *Queue[T]) IsEmpty() bool {
	return (len(q.items) == 0)
}

func (q *Queue[T]) Clear() {
	q.items = q.items[:0]
}

func (q *Queue[T]) Print() {
	if q.IsEmpty() {
		fmt.Println(q.items)
		return
	}
	for _, item := range q.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
