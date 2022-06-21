package stack

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (stack *Stack) Push(d int) {
	newNode := &Node{data: d}
	if stack.top != nil {
		newNode.next = stack.top
	}
	stack.top = newNode
}

func (stack *Stack) Pop() (int, bool) {
	if stack.IsEmpty() {
		return 0, false
	} else {
		data := stack.top.data
		stack.top = stack.top.next
		return data, true
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}

func (stack *Stack) Peek() (int, bool) {
	if stack.IsEmpty() {
		return 0, false
	} else {
		return stack.top.data, true
	}
}

func (stack *Stack) Print() {

	if stack.IsEmpty() {
		fmt.Println("Stack is empty.")
	} else {
		fmt.Print("Printing stack: ")
	}

	for temp := stack.top; temp != nil; temp = temp.next {
		fmt.Printf("%v ", temp.data)
	}
}
