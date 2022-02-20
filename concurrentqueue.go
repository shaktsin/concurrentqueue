package concurrentqueue

import (
	"fmt"
	"sync"
)

type ConcurrentQueue struct {
	q  *Queue
	mu sync.Mutex
}

func NewConcurrentQueue(cap int) *ConcurrentQueue {
	return &ConcurrentQueue{
		q: NewQueue(cap),
	}
}

func (q *ConcurrentQueue) String() string {
	return q.q.String()
}

func (q *ConcurrentQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.q.Size()
}

func (q *ConcurrentQueue) Inqueue(val interface{}) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	fmt.Printf("Queue snapshot %s\n", q.q)

	return q.q.Inqueue(val)
}

func (q *ConcurrentQueue) Dequeue() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	fmt.Printf("Queue snapshot %s\n", q.q)
	return q.q.Dequeue()
}
