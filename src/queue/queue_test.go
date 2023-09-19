package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	input := []int{5, 734445, 0, -1}

	q := New[int]()
	for _, i := range input {
		q.Push(i)
	}

	o, _ := q.Peek()
	if o != input[0] {
		t.Errorf("Peek failed: Got: %v, Expected: %v\n", o, input[0])
	}

	for _, i := range input {
		o, _ := q.Pop()
		if i != o {
			t.Errorf("Pop failed: Got: %v, Expected: %v, Index %v \n", o, input[i], i)
		}
	}

	if q.Len() != 0 {
		t.Errorf("Length was %v \n", q.Len())
	}
}
