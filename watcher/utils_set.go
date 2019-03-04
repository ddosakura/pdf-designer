package watcher

// TODO: 之后处理一下一下（单独提出去之类的）

import "sync"

// Set ...
type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

// NewSet ...
func NewSet() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

// Add ...
func (s *Set) Add(item interface{}) {
	//写锁
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

// Remove ...
func (s *Set) Remove(item interface{}) {
	//写锁
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// Has ...
func (s *Set) Has(item interface{}) bool {
	//允许读
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// List ...
func (s *Set) List() []interface{} {
	//允许读
	s.RLock()
	defer s.RUnlock()
	var outList []interface{}
	for value := range s.m {
		outList = append(outList, value)
	}
	return outList
}

// Len ...
func (s *Set) Len() int {
	return len(s.List())
}

// Clear ...
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

// IsEmpty ...
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}
