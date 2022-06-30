package linked_list

import (
	"testing"
)

func TestReverseOne(t *testing.T) {
	ll := &LinkedList{}
	ll.InsertAt(17, 0)
	ll.InsertAt(1, 1)
	ll.InsertAt(98, 2)
	ll.InsertAt(18, 3)
	ll.InsertAt(5, 4)

	s1 := "Printing list: 5 18 98 1 17\n"
	ll.head = ReverseOne(ll.Head())
	s2 := ll.Print()

	if s1 != s2 {
		t.Error("Expected", s1, "Got", s2)
	} else {
		t.Log("PASSED.")
	}
}