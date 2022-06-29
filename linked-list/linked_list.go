package linked_list

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (ll *LinkedList) Len() int {
	return ll.len
}

func (ll *LinkedList) Head() *Node {
	return ll.head
}

func (ll *LinkedList) Tail() *Node {
	return ll.tail
}

func (ll *LinkedList) InsertAt(dt, index int) bool {

	newNode := &Node{data: dt}

	if index > ll.len || index < 0 {
		return false
	} else if index == 0 {
		ll.insertAtHead(newNode)
	} else if index == ll.len {
		ll.insertAtTail(newNode)
	} else {
		temp := ll.head
		for i := 0; i < index-1; i++ {
			temp = temp.next
		}
		newNode.next = temp.next
		temp.next = newNode
	}
	ll.len++

	return true
}

func (ll *LinkedList) insertAtHead(node *Node) {
	node.next = ll.head
	ll.head = node
	ll.tail = node
}

func (ll *LinkedList) insertAtTail(node *Node) {
	ll.tail.next = node
	ll.tail = node
}

func (ll *LinkedList) DataAt(index int) (int, bool) {
	if ll.len <= 0 || index >= ll.len {
		return 0, false
	}

	temp := ll.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.data, true
}

func (ll *LinkedList) Print() {
	if ll.len == 0 {
		fmt.Println("List is empty.")
		return
	} else {
		fmt.Print("Printing list:")
	}

	for temp := ll.head; temp != nil; temp = temp.next {
		fmt.Print(" ", temp.data)
	}
	fmt.Println()
}

// Find returns the index of the dataToFind and whether is present in list or not. 
func (ll *LinkedList) Find(dataToFind int) (int, bool) {
	for i, temp := 0, ll.head; i < ll.len; i, temp = i+1, temp.next {
		if temp.data == dataToFind {
			return i, true
		}
	}
	return 0, false
}
