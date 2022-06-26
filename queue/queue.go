package queue

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node // pointer to the first node
	rear  *Node // pointer to the last node
}

func (queue *Queue) IsEmpty() bool {
	return queue.front == nil && queue.rear == nil
}

func (queue *Queue) Enqueue(d int) {
	newNode := &Node{data: d}

	if queue.IsEmpty() {
		queue.front = newNode
		queue.rear = newNode
	} else {
		queue.rear.next = newNode
		queue.rear = newNode
	}
}

func (queue *Queue) Dequeue() (int, bool) {
	if queue.IsEmpty() {
		return 0, false
	} else {
		data := queue.front.data
		queue.front = queue.front.next
		if queue.front == nil {
			queue.rear = nil
		}
		return data, true
	}
}

func (queue *Queue) Peek() (int, bool) {
	if queue.IsEmpty() {
		return 0, false
	} else {
		return queue.front.data, true
	}
}
