package deque

import "log"

type Deque struct {
	front int
	rear  int
	size  int
	queue []any
}

func New(size int) *Deque {
	return &Deque{
		front: -1,
		rear:  0,
		size:  size,
		queue: make([]any, size),
	}
}

func (d *Deque) InsertFront(element any) {
	if d.IsFull() {
		log.Println("insert: deque is full")
		return
	}

	if d.front < 1 {
		d.front = d.size - 1
	}

	d.front--
	d.queue[d.front] = element
}

func (d *Deque) InsertRear(element any) {
	if d.IsFull() {
		log.Println("insert: deque is full")
		return
	}

	if d.rear == d.size-1 {
		d.rear = 0
	}

	d.rear++
	d.queue[d.rear] = element
}

func (d *Deque) DeleteFront() {
	if d.IsEmpty() {
		log.Println("delete: deque is empty")
		return
	}

	d.queue[d.front] = nil

	// check for last element
	if d.front == d.rear {
		d.front, d.rear = -1, -1
		return
	}

	// check for front being at the end
	if d.front == d.size-1 {
		d.front = 0
		return
	}

	d.front++
}

func (d *Deque) DeleteRear() {
	if d.IsEmpty() {
		log.Println("delete: deque is empty")
		return
	}

	d.queue[d.rear] = nil

	// check for last element
	if d.rear == d.front {
		d.front, d.rear = -1, -1
	}

	// check for position of rear
	if d.rear == 0 {
		d.rear = d.size - 1
	}
}

func (d *Deque) IsFull() bool {
	if d.front == d.rear+1 || d.front == 0 && d.rear == d.size-1 {
		return true
	}

	return false
}

func (d *Deque) IsEmpty() bool {
	return d.front == -1
}
