package equal

import (
	"testing"
)

func TestIsCycleStructCase(t *testing.T) {
	type Cycle struct {
		value int
		tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}

	if !IsCycle(c) {
		t.Error("IsCycle() got false, but want true")
	}
}

func TestIsRecursiveSliceCase(t *testing.T) {
	type S []S
	var s = make(S, 1)
	s[0] = s

	if !IsCycle(s) {
		t.Error("IsCycle(s) got false, but want true")
	}
}

func TestIsCycleSliceCase(t *testing.T) {
	type slice struct {
		next []*slice
	}

	a := &slice{}
	b := &slice{}
	c := &slice{}

	a.next = append(a.next, b)
	b.next = append(b.next, c)
	c.next = append(c.next, a)

	if !IsCycle(a) {
		t.Error("IsCycle(a) got false, but want true")
	}
}

func TestIsCycleMapCase(t *testing.T) {
	type cycleMap struct {
		next map[string]*cycleMap
	}

	a := &cycleMap{make(map[string]*cycleMap)}
	b := &cycleMap{make(map[string]*cycleMap)}
	c := &cycleMap{make(map[string]*cycleMap)}

	a.next["a"] = b
	b.next["b"] = c
	c.next["c"] = a

	if !IsCycle(a) {
		t.Error("IsCycle(a) got false, but want true")
	}
}
