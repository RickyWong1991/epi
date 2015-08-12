// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

import "testing"

// createCycleList returns cycled linked list created from data.
// The cycle in the list is created to node on index ci of data.
// The returned node represents the reference to the start of the
// cycle or nil is returned when there is no cycle in returned list.
// If ci < 0 || ci >= len(data), then a list without cycle is returned.
func createCycleList(data []interface{}, ci int) (*List, *Node) {
	if ci < 0 || ci >= len(data) {
		return NewFromSlice(data), nil
	}

	l := new(List)
	var csn *Node
	for i := 0; i <= ci; i++ {
		csn = l.Insert(&Node{Data: data[i]})
	}
	var ln *Node
	for i := ci; i < len(data); i++ {
		ln = l.Insert(&Node{Data: data[i]})
	}
	ln.next = csn // Create a cycle from the ln to the csn.

	return l, csn
}

func TestHasCycle(t *testing.T) {
	for _, test := range []struct {
		l []interface{}
		i int
	}{
		// Test empty list.
		{[]interface{}{}, -1},

		// Test lists without cycle.
		{[]interface{}{0}, -1},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, -1},

		// Test lists with cycle.
		{[]interface{}{0}, 0},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, 0},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, 3},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, 7},
	} {
		l, want := createCycleList(test.l, test.i)
		if got := HasCycle(l); got != want {
			t.Errorf("HasCycle(%v) = %v; want %v", test.l, got, want)
		}
	}
}

func benchHasCycle(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]interface{}, size) // We don't care about content but about pointers.
		l, n := createCycleList(data, 0)
		b.StartTimer()
		csn := HasCycle(l)
		b.StopTimer()
		if n != csn {
			b.Error("HasCycle did not find the cycle")
		}
	}
}

func BenchmarkHasCycle1e1(b *testing.B) { benchHasCycle(b, 1e1) }
func BenchmarkHasCycle1e2(b *testing.B) { benchHasCycle(b, 1e2) }
func BenchmarkHasCycle1e3(b *testing.B) { benchHasCycle(b, 1e3) }
