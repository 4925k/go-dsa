package queue

import (
	"fmt"
	"log"
)

type Queue struct {
	front int
	rear  int
	size  int
	queue []any
}

func New(size int) *Queue {
	return &Queue{
		front: -1,
		rear:  -1,
		size:  size,
		queue: make([]any, size),
	}
}

func (q *Queue) Enqueue(element any) {
	// check if queue is full
	if q.rear == q.size-1 {
		log.Println("enqueue: queue is full")
		return
	}

	// first element
	if q.rear == -1 {
		q.front = 0
	}

	q.rear++
	q.queue[q.rear] = element
}

func (q *Queue) Dequeue() any {
	// check for empty queue
	if q.front == -1 {
		log.Println("dequeue: queue is empty")
		return nil
	}

	// check if last element
	if q.front == q.size-1 {
		buffer := q.queue[q.front]
		q.queue[q.front] = nil
		q.front, q.rear = -1, -1
		return buffer
	}

	buffer := q.queue[q.front]
	q.queue[q.front] = nil
	q.front++

	return buffer
}

func (q *Queue) Display() {
	fmt.Println(q.queue)
}

func (q *Queue) Size() int {
	return q.size
}

func QueueTest() {
	q := New(3)

	// display
	q.Display()

	// adding
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	// display
	q.Display()

	// removing
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	// display
	q.Display()
}
