package stack

import "log"

type Stack []any

// Create will return a Stack
func Create(size int) Stack {
	return make([]any, 0, size)
}

// Len returns current size of the Stack
func (s Stack) Len() int {
	return len(s)
}

// Cap returns the size of the Stack
func (s Stack) Cap() int {
	return cap(s)
}

// Empty check if Stack empty
func (s Stack) Empty() bool {
	return s.Len() == 0
}

// Full checks if the Stack if full
func (s Stack) Full() bool {
	return s.Len() == s.Cap()
}

// Push will add items to the Stack
func (s *Stack) Add(item any) {
	if s.Full() {
		log.Println("add failed: Stack if full")
		return
	}

	*s = append(*s, item)
}

// Pop will remove an element from the Stack
// returns empty if no element
func (s *Stack) Pop() any {
	if s.Empty() {
		log.Println("pop failed: Stack is empty")
		return ""
	}

	buffer := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return buffer
}
