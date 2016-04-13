package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack(10)

	for i := 1; i <= 10; i++ {
		err := s.Push(i)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestPop(t *testing.T) {
	s := NewStack(10)

	for i := 1; i <= 10; i++ {
		s.Push(i)
	}

	for i := 10; i > 0; i-- {
		x, err := s.Pop()
		if err != nil {
			t.Error(err)
		}
		if x != i {
			t.Errorf("expect %d, actual %d", x, i)
		}
	}
}
