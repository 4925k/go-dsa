package binarytree

import "fmt"

type Node struct {
	data  int
	right *Node
	left  *Node
}

type Tree struct {
	root *Node
}

func NewNode(data int) *Node {
	return &Node{
		data:  data,
		right: nil,
		left:  nil,
	}
}

func (n *Node) DepthFirstTraversal() []int {
	if n == nil {
		return nil
	}

	left := n.left.DepthFirstTraversal()
	right := n.right.DepthFirstTraversal()

	res := []int{n.data}

	return append(res, append(left, right...)...)

}

func (n *Node) BreadthFirstTraversal() []int {
	if n == nil {
		return nil
	}

	s := []*Node{n} // act as a queue
	res := []int{}  // store values
	for len(s) > 0 {

		currentNode := s[0]
		res = append(res, currentNode.data)
		s = s[1:]

		if currentNode.left != nil {
			s = append(s, currentNode.left)
		}

		if currentNode.right != nil {
			s = append(s, currentNode.right)
		}
	}

	return res
}

func (n *Node) DepthFirstSearch(target int) bool {
	if n == nil {
		return false
	}

	if n.data == target {
		return true
	}

	if n.left.DepthFirstSearch(target) || n.right.DepthFirstSearch(target) {
		return true
	}

	return false
}

func (n *Node) BreadthFirstSearch(target int) bool {
	if n == nil {
		return false
	}

	s := []*Node{n} // act as a queue for BFT

	for len(s) > 0 {
		currentNode := s[0]
		s = s[1:]

		if currentNode.data == target {
			return true
		}

		if currentNode.left != nil {
			s = append(s, currentNode.left)
		}

		if currentNode.right != nil {
			s = append(s, currentNode.right)
		}

	}

	return false
}

func (n *Node) SumDepthFirst() int {
	if n == nil {
		return 0
	}

	return n.data + n.left.SumDepthFirst() + n.right.SumDepthFirst()
}

func (n *Node) SumBreadthFirst() int {
	if n == nil {
		return 0
	}

	var sum int
	queue := []*Node{n}

	for len(queue) > 0 {
		currentNode := queue[0]
		sum += currentNode.data
		queue = queue[1:]

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}

		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}

	return sum

}

func BtTest() {
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)
	d := NewNode(4)
	e := NewNode(5)
	f := NewNode(6)

	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f

	// test functions
	fmt.Println("Depth First Travesal\t", a.DepthFirstTraversal())
	fmt.Println("Breadth First Travesal\t", a.BreadthFirstTraversal())

	fmt.Println("Depth First Search\t", a.DepthFirstSearch(12))
	fmt.Println("Breadth First Search\t", a.BreadthFirstSearch(111))

	fmt.Println("Depth First Sum\t\t", a.SumDepthFirst())
	fmt.Println("Breadth First Sum\t", a.SumBreadthFirst())

}
