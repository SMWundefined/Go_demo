package main

import "fmt"

var count int = 0

//Node declaration
type Node struct {
	Key   int
	left  *Node
	right *Node
}

//Insert adding nodes to the tree
//The k here is different from the root Key
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.right == nil {
			n.right = &Node{Key: k}
		} else {
			n.right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.left == nil {
			n.left = &Node{Key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

//Search will take a key value and return true if the value is in the tree

func (n *Node) Search(s int) bool {

	count++

	if n == nil {
		return false
	}
	if n.Key < s {
		return n.right.Search(s)
	} else if n.Key > s {
		return n.left.Search(s)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(54)
	tree.Insert(76)
	tree.Insert(500)
	tree.Insert(44)
	tree.Insert(33)
	tree.Insert(69)
	tree.Insert(4)

	fmt.Println(tree.Search(54))
	fmt.Println(count)
	//fmt.Println(tree)
}
