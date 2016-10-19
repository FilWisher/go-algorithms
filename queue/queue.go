package queue

type QueueItem struct {
	Val interface{}
	Next *QueueItem
}

type Queue struct {
	len int
	Front, Back *QueueItem
}

func NewQueue() *Queue {
	return &Queue{
		len: 0,
		Front: nil,
		Back: nil,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue) Enqueue(val interface{}) {
	old := q.Back
	q.Back = &QueueItem{Val:val}
	if q.IsEmpty() {
		q.Front = q.Back
	} else {
		old.Next = q.Back
	}
	q.len++
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	val := q.Front.Val
	q.Front = q.Front.Next
	if q.IsEmpty() {
		q.Back = q.Front
	}
	q.len--
	return val, true
}

func (q *Queue) Each() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		item, ok := q.Dequeue()
		for ; !q.IsEmpty() && ok; item, ok = q.Dequeue() {
			ch <- item
		}
		close(ch)
	}()
	return ch
}

