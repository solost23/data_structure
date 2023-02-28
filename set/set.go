package set

import "sync"

type Set struct {
	set map[any]struct{}
	len int
	sync.RWMutex
}

func NewSet(cap int) *Set {
	return &Set{
		set: make(map[any]struct{}, cap),
	}
}

func (s *Set) Add(elem ...any) {
	s.Lock()
	defer s.Unlock()

	if len(elem) <= 0 {
		return
	}
	for _, e := range elem {
		s.set[e] = struct{}{}
	}
	s.len = len(s.set)
}

func (s *Set) Remove(elem any) {
	s.Lock()
	defer s.Unlock()

	if s.len <= 0 {
		return
	}

	delete(s.set, elem)
	s.len = len(s.set)
}

func (s *Set) Has(elem any) bool {
	s.RLock()
	defer s.RUnlock()

	if s.len <= 0 {
		return false
	}
	_, ok := s.set[elem]
	if !ok {
		return false
	}
	return true
}

func (s *Set) Len() int {
	s.RLock()
	defer s.RUnlock()

	return s.len
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()

	s.set = map[any]struct{}{}
	s.len = 0
}

func (s *Set) IsEmpty() bool {
	if s.Len() <= 0 {
		return true
	}
	return false
}
