package queue

type Queue[T any] struct {
	arr []T

	len  int
	head int
	tail int
}

func (q *Queue[T]) Push(t T) {
	q.len++
	if q.len >= len(q.arr) {
		narr := make([]T, (len(q.arr)+1)*2)
		copy(narr, q.arr)
		q.arr = narr

		q.head = q.len - 1
		q.tail = 0
	}

	q.arr[q.head] = t
	q.head = (q.head + 1) % len(q.arr)
}

func (q *Queue[T]) Pop() (T, bool) {
	var t T
	if q.len < 1 {
		return t, false
	}
	t = q.arr[q.tail]
	q.tail = q.tail + 1%len(q.arr)
	q.len--

	return t, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var t T
	if q.len < 1 {
		return t, false
	}

	t = q.arr[q.tail]

	return t, true
}

func (q *Queue[T]) Len() int {
	return q.len
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}
