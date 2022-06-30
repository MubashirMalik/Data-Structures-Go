package linked_list

func ReverseOne(head *Node) *Node {
	var t1, t2 *Node
	for head != nil {
		t2 = head.next
		head.next = t1
		t1 = head
		head = t2
	}
	return t1
}

/* CHALLENGE # 01

Given the pointer to the head node of a linked list, change the next pointers of the nodes so that their order is reversed. The head pointer given may be null meaning that the initial list is empty.

*/
