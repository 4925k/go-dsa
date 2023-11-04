package singly

import (
	"fmt"
	"log"
)

type Singly struct {
	head *List
	tail *List
}

type List struct {
	data int
	prev *List
	next *List
}

func New(data int) *Singly {
	new := &List{
		data: data,
	}

	return &Singly{
		head: new,
		tail: new,
	}
}

func (s *Singly) InsertFront(data int) {
	new := &List{
		data: data,
		next: s.head,
	}

	s.head.prev = new
	s.head = new
}

func (s *Singly) InsertEnd(data int) {
	new := &List{
		data: data,
		prev: s.tail,
	}

	s.tail.next = new
	s.tail = new
}

func (s *Singly) InsertIndex(n int, data int) {
	pointer := s.head
	for i := 0; i <= n; i++ {
		if pointer == nil {
			return
		}

		if i == n {
			new := &List{
				data: data,
				prev: pointer.prev,
				next: pointer,
			}

			pointer.prev.next = new
			pointer.prev = new
		}

		pointer = pointer.next
	}
}

func (s *Singly) DeleteFront() {
	s.head.next.prev = nil
	s.head = s.head.next
}

func (s *Singly) DeleteLast() {
	s.tail.prev.next = nil
	s.tail = s.tail.prev
}

func (s *Singly) DeleteIndex(n int) {
	pointer := s.head

	for i := 0; i <= n; i++ {
		if pointer == nil {
			log.Println("delete: invalid position")
			return
		}

		if i == n {
			pointer.prev.next = pointer.next
			pointer.next.prev = pointer.prev
		}

		pointer = pointer.next
	}
}

func (s *Singly) Search(data int) bool {
	pointer := s.head

	for pointer != nil {
		if pointer.data == data {
			return true
		}

		pointer = pointer.next
	}

	return false
}

func (s *Singly) Display() {
	pointer := s.head
	for pointer != nil {
		fmt.Println(pointer.data)
		pointer = pointer.next
	}
}

func (s *Singly) Sort() {
	for pointer := s.head; pointer != nil; pointer = pointer.next {
		for pointer2 := pointer.next; pointer2 != nil; pointer2 = pointer2.next {
			if pointer.data > pointer2.data {
				pointer.data, pointer2.data = pointer2.data, pointer.data
			}
		}

	}
}

func SinglyTest() {
	list := New(5)

	fmt.Println("new")
	list.Display()

	list.InsertFront(6)
	list.InsertEnd(4)
	list.InsertEnd(8)
	list.InsertEnd(9)
	list.InsertIndex(1, 10)

	fmt.Println("insert")
	list.Display()

	list.DeleteFront()
	list.DeleteLast()
	list.DeleteIndex(2)

	fmt.Println("delete")
	list.Display()

	list.Sort()

	fmt.Println("sort")
	list.Display()

}
