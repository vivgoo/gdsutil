package gdsutil

type Node struct {
	data interface{}
	next *Node
}
type Queue struct {
	head      *Node
	tail      *Node
	itemCount int
}

func NewQueue() *Queue {
	return &Queue{
		nil,
		nil,
		0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.itemCount == 0
}

func (q *Queue) EnQueue(data interface{}) {
	node := &Node{
		data,
		nil,
	}

	if q.itemCount == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.itemCount++
}

func (q *Queue) DeQueue() interface{} {
	if 0 == q.itemCount {
		return nil
	}

	node := q.head
	q.head = q.head.next
	q.itemCount--
	if 0 == q.itemCount {
		q.tail = nil
	}
	return node.data
}

func (q *Queue) HeadNode() interface{} {
	if 0 == q.itemCount {
		return nil
	}
	return q.head.data
}
