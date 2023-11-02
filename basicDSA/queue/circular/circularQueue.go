package circularQueue

import (
	"fmt"
	"log"
)

type CircularQueue struct {
	front   int
	rear    int
	pointer int
	size    int
	queue   []any
}

func New(size int) *CircularQueue {
	return &CircularQueue{
		front:   -1,
		rear:    -1,
		pointer: -1,
		size:    size,
		queue:   make([]any, size),
	}
}

// Enqueue adds an element to the queue
func (q *CircularQueue) Enqueue(element any) {
	// check if queue is full
	if q.IsFull() {
		log.Println("enqueue: queue is full")
		return
	}

	// set front to zero if first value
	if q.front == -1 {
		q.front = 0
	}

	// increase rear pointer, if exceeeds size, go to first
	if q.rear++; q.rear >= q.size {
		q.rear = 0
	}

	// set element
	q.queue[q.rear] = element
}

// Dequeue pops the element in the front of the queue
func (q *CircularQueue) Dequeue() any {
	// check for empty queue
	if q.IsEmpty() {
		log.Println("dequeue: queue is empty")
		return nil
	}

	buffer := q.queue[q.front] // store value

	// check for last element
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
		return buffer
	}

	// increase front value
	q.front = (q.front + 1) % q.size

	return buffer
}

// IsFull checks if the queue if full
func (q *CircularQueue) IsFull() bool {
	if q.front == q.rear+1 || q.front == 0 && q.rear == q.size-1 {
		return true
	}

	return false
}

// IsEmpty checks if the queue if empty
func (q *CircularQueue) IsEmpty() bool {
	return q.front == -1
}

func (q *CircularQueue) Display() {
	fmt.Println(q.queue)
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

	// removing and adding element in the first slot
	fmt.Println(q.Dequeue())
	q.Enqueue(4)

	q.Display()

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	// display
	q.Display()
}
