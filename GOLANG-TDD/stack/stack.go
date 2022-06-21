package stack

import "fmt"

var ErrNoSuchElement = fmt.Errorf("NO SUCH ELEMENT")

func New() *Stack {
	return &Stack{
		arr: make([]int, 0),
	}
}

type Stack struct{
	arr []int
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Push(val int) error {
	s.arr = append(s.arr, val)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNoSuchElement
	}
	value := s.getLastElement()
	s.arr = s.arr[:s.Size() - 1]
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNoSuchElement
	}
	return s.getLastElement(), nil
}

func (s Stack) getLastElement() int {
	return s.arr[s.Size() - 1]
}
