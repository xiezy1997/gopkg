package gstack

type stack[t any] struct {
	items []t
	len   int
}

// The NewStack function in Go creates a new stack with a specified type.
func NewStack[t any]() IStahk[t] {
	return &stack[t]{make([]t, 0), 0}
}

func (s *stack[t]) IsEmpty() bool { return s.Size() == 0 }

func (s *stack[t]) Size() int { return s.len }

func (s *stack[t]) Clear() {
	s.items = s.items[:0]
	s.len = 0
}

func (s *stack[t]) TPush(item t) {
	s.items = append(s.items, item)
	s.len += 1
}

func (s *stack[t]) TPop() (t, bool) {
	if s.IsEmpty() {
		var tmp t
		return tmp, false
	}
	i := s.items[s.len-1]
	s.len -= 1
	s.items = s.items[:s.len]
	return i, true
}
