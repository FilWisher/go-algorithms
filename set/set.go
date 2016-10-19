package set

type Set struct {
	set map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{
		set: make(map[interface{}]struct{}),	
	}
}

func (s *Set) Add(it interface{}) {
	s.set[it] = struct{}{}
}

func (s *Set) Contains(it interface{}) bool {
	_, ok := s.set[it]
	return ok
}

func (s *Set) Each() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for k := range s.set {
			ch <- k
		}	
		close(ch)
	}()
	return ch
}
