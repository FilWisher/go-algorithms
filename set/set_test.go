package set_test

import (
	"testing"
	"github.com/filwisher/algo/set"
)

func TestSet(t *testing.T) {
	s := set.NewSet()
	
	s.Add(2)
	
	if !s.Contains(2) {
		t.Errorf("set does not contain added element")
	}
	
	for i := 3; i < 10; i++ {
		if s.Contains(i) {
			t.Errorf("set contains unadded elements")
		}
	}
	
	s.Add(2)

	i := 0
	for v := range s.Each() {
		i += 1	
		if v != 2 {
			t.Errorf("set should only contain 2")
		}
	}
	if i != 1 {
		t.Errorf("set contains duplicates")
	}	
}
