package main

import "fmt"

type node struct{
	data int
	next *node
}
type linkedlist struct{
	head *node
	len int
}

func (l *linkedlist) prepend (n *node){
	secondElement := l.head
	l.head=n
	l.head.next=secondElement
	l.len++ 

}

func main(){
	//start  := time.Now()
	mylist := linkedlist{}
	node1 := &node{data:55}
	node2 := &node{data:1}
	node3 := &node{data:33}
	node4 := &node{data:21}
	node5 := &node{data:9}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	fmt.Println(mylist)
	// mylist.display()
	// mylist.delete(900)
	// mylist.display()
	// emptyList :=linkedlist{}
	// emptyList.delete(90)

}