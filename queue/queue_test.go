package queue_test

import (
	"github.com/filwisher/algo/queue"
	"testing"
)

func TestQueue(t *testing.T) {
	q := queue.NewQueue()
	if !q.IsEmpty() {
		t.Errorf("new queue not empty")
	}
	
	for i := 0; i < 10; i++ {
		q.Enqueue(i)	
	}
	
	if q.IsEmpty() {
		t.Errorf("queue not empty after being filled")
	}

	for i := 0; i < 10; i++ {
		x, ok := q.Dequeue()
		if !ok {
			t.Errorf("queue is empty, not enough items stored (%d%)", i)
		}
		if x != i {
			t.Errorf("queue does not maintain correct order")
		}
	}
	
	if !q.IsEmpty() {
		t.Errorf("IsEmpty false after queue is emptied")
	}
}

func TestQueueEach(t *testing.T) {
	q := queue.NewQueue()
	
	for i := 0; i < 10; i++ {
		q.Enqueue(i)	
	}
	
	i := 0
	for x := range q.Each() {
		if x != i {
			t.Errorf("queue items not returned in correct order")
		}
		i++
	}
}
