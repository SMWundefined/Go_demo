package main

import (
	"fmt"
)


type node struct{
	data int
	next *node
}

type linkedlist struct{
	head *node
	len int
}

func (l *linkedlist)prepend(n *node){
	second :=l.head
	l.head=n
	l.head.next=second
	l.len++
}

func (l linkedlist) display(){
	displayList := l.head
	for l.len !=0{
		fmt.Printf("%d\t",displayList.data)
		displayList = displayList.next
		l.len--
	}
	fmt.Printf("\n")
}

func (l *linkedlist) delete(value int){
	if l.len==0{
		return
	}
	if l.head.data==value{
		l.head=l.head.next
		l.len--
		return
	}
	deleteElement := l.head
	for deleteElement.next.data !=value{
		if deleteElement.next.next==nil{
			return
		}
		deleteElement=deleteElement.next
		//update for loop here^
	}
	deleteElement.next=deleteElement.next.next
	fmt.Printf("Deletion of %d is complete\n", value)
	l.len--
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
	mylist.display()
	mylist.delete(900)
	mylist.display()
	emptyList :=linkedlist{}
	emptyList.delete(90)

}