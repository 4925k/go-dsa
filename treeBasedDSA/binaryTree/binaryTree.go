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

func BtTest() {
	a := NewNode(1)
	b := NewNode(5)
	c := NewNode(12)
	d := NewNode(13)
	e := NewNode(14)
	f := NewNode(15)

	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f

	fmt.Println(a.DepthFirstTraversal())
}
