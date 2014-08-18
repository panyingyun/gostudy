package stacker

import (
	"errors"
)

type Stack struct {
	st  []interface{}
	len int
	cap int
}

func NewStack(cap int) *Stack {
	st := make([]interface{}, 0, cap)
	return &Stack{st, 0, cap}
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Cap() int {
	return s.cap
}

func (s *Stack) Push(p interface{}) {
	s.st = append(s.st, p)
	s.len = len(s.st)
	s.cap = cap(s.st)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.len == 0 {
		return nil, errors.New("Can't pop an empty stack")
	}
	s.len -= 1
	out := s.st[s.len]
	s.st = s.st[:s.len]
	return out, nil
}
