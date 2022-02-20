package concurrentqueue

import (
	"fmt"
	"strings"
)

// node in a queue
type Node struct {
	prev *Node
	next *Node
	val  interface{}
}

// create node operation
func newNode(val interface{}) *Node {
	return &Node{
		prev: nil,
		next: nil,
		val:  val,
	}
}

// queue definition
type Queue struct {
	first *Node
	last  *Node
	size  int // size of the queue
	cap   int // max capacity a queue can handle
}

// create new queue
func NewQueue(cap int) *Queue {
	return &Queue{
		cap: cap,
	}
}

func (q *Queue) String() string {

	var sb strings.Builder
	temp := q.first

	if temp == nil {
		return ""
	}

	for temp != nil {
		sb.WriteString(fmt.Sprintf("%v", temp.val))
		sb.WriteString("->")
		temp = temp.next
	}
	return sb.String()[:len(sb.String())-2]
}

// size of the queue
func (q *Queue) Size() int {
	return q.size
}

// inqueue operation
func (q *Queue) Inqueue(val interface{}) error {
	// capacity check
	if q.Size() == q.cap {
		return fmt.Errorf("Queue overflow: capacity %d", q.cap)
	}

	node := newNode(val)
	if q.Size() == 0 {
		q.first = node
		q.last = node
	} else {
		// set bidirectional pointers
		q.last.next = node
		node.prev = q.last

		// update last node of queue
		q.last = node
	}

	// increment size of the queue
	q.size++
	return nil
}

// dequeue operation
func (q *Queue) Dequeue() (interface{}, error) {
	// underflow check
	if q.Size() == 0 {
		return nil, fmt.Errorf("Queue underflow")
	}

	last := q.last
	if q.Size() == 1 {
		q.first = nil
		q.last = nil
	} else {
		prevlast := q.last.prev
		prevlast.next = nil
		q.last = prevlast
	}

	// decrement size
	q.size--

	return last.val, nil
}
