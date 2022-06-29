package linked_list

import (
	"testing"
)

func TestInsertAt(t *testing.T) {
	ll := LinkedList{}

	type test struct {
		dataIn  int
		indexIn int
		result  bool
	}
	tests := []test{
		{42, -1, false},
		{42, 1, false},
		{42, 0, true},
		{43, 1, true},
		{46, 2, true},
		{45, 2, true},
		{44, 2, true},
	}

	for _, v := range tests {
		result := ll.InsertAt(v.dataIn, v.indexIn)
		if v.result == result {
			t.Log("PASSED.")
		} else {
			t.Error("Expected", v.result, "Got", result)
		}
	}
}

func TestDataAt(t *testing.T) {
	ll := LinkedList{}

	ll.InsertAt(42, 0)
	if data, _ := ll.DataAt(0); data != 42 {
		t.Error("Expected", 42, "Got", data)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(43, 1)
	if data, _ := ll.DataAt(1); data != 43 {
		t.Error("Expected", 43, "Got", data)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(46, 2)
	if data, _ := ll.DataAt(2); data != 46 {
		t.Error("Expected", 46, "Got", data)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(45, 2)
	if data, _ := ll.DataAt(2); data != 45 {
		t.Error("Expected", 45, "Got", data)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(44, 2)
	if data, _ := ll.DataAt(2); data != 44 {
		t.Error("Expected", 44, "Got", data)
	} else {
		t.Log("PASSED.")
	}
}

func TestFind(t *testing.T) {
	ll := LinkedList{}

	ll.InsertAt(42, 0)
	if index, _ := ll.Find(42); index != 0 {
		t.Error("Expected", 0, "Got", index)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(43, 1)
	if index, _ := ll.Find(43); index != 1 {
		t.Error("Expected", 1, "Got", index)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(46, 2)
	if index, _ := ll.Find(46); index != 2 {
		t.Error("Expected", 2, "Got", index)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(45, 2)
	if index, _ := ll.Find(46); index != 3 {
		t.Error("Expected", 3, "Got", index)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(44, 2)
	if index, _ := ll.Find(44); index != 2 {
		t.Error("Expected", 2, "Got", index)
	} else {
		t.Log("PASSED.")
	}

	ll.InsertAt(44, 4)
	if index, _ := ll.Find(44); index != 2 {
		t.Error("Expected", 2, "Got", index)
	} else {
		t.Log("PASSED.")
	}
}
