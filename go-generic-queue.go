package gogenericqueue

type Queue[T interface{}] struct {
	data     []T
	capacity uint64
}

func QueueInit[T interface{}]() *Queue[T] {
	return &Queue[T]{
		data:     make([]T, 0),
		capacity: 0,
	}
}

func (q *Queue[T]) Push(elem T) {
	q.data = append(q.data, elem)
	q.capacity += 1
}

func (q *Queue[T]) Pop() (elem T) {
	if q.capacity == 0 {
		return
	}
	tmp := q.data[0]
	q.data = q.data[1:]
	q.capacity -= 1
	return tmp
}

func (q *Queue[T]) Cap() uint64 {
	return q.capacity
}
