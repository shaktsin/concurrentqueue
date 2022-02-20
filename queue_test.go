package concurrentqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInQueueEmpty(t *testing.T) {
	q := NewQueue(0)
	err := q.Inqueue(1)
	assert.NotNil(t, err)
}

func TestInqueue(t *testing.T) {
	q := NewQueue(2)
	q.Inqueue(1)
	q.Inqueue(2)

	val, _ := q.Dequeue()
	assert.Equal(t, 2, val)

	val, _ = q.Dequeue()
	assert.Equal(t, 1, val)

	val, err := q.Dequeue()
	assert.NotNil(t, err)
}
