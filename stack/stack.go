package stack

import (
	"sync"
)

type Stack struct {
	stack []any
	tt    int
	sync.RWMutex
}

func NewStack(size int) *Stack {
	return &Stack{
		stack: make([]any, size),
		tt:    -1,
	}
}

func (s *Stack) Push(elem ...any) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	if len(elem) <= 0 {
		return
	}
	for _, e := range elem {
		s.tt++
		s.stack[s.tt] = e
	}
}

func (s *Stack) Query() any {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	return s.stack[s.tt]
}

func (s *Stack) Pop() {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	s.tt--
}
