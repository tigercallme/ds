package ring

import (
	"testing"
)

func TestCapacity(t *testing.T) {
	q := NewRingQueue(10)

	if q.capacity != 10 {
		t.Fatal("Capacity of ring was not 10", q.capacity)
	}
}

func TestPush(t *testing.T) {
	q := NewRingQueue(10)

	for i := 1; i <= 10; i++ {
		err := q.Push(i)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestPop(t *testing.T) {
	q := NewRingQueue(5)

	var tests = []struct {
		i      int
		expect int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
	}

	for _, tt := range tests {
		q.Push(tt.i)
		actual, err := q.Pop()

		if err != nil {
			t.Error(err)
		}
		if actual != tt.expect {
			t.Errorf("push %d, expect %d, actual %d", tt.i, tt.expect, actual)
		}
	}

}
