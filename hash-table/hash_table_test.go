package hashtable

import (
	"math/rand"
	"testing"
	"time"
)

func TestSetHashGroups(t *testing.T) {
	var ht HashTable
	if ht.SetHashGroups(10) == true && ht.GetHashGroups() == 10 {
		t.Log("PASSED.")
	} else {
		t.Error("FAILED")
	}

	if ht.SetHashGroups(11) == false && ht.GetHashGroups() == 10 {
		t.Log("PASSED.")
	} else {
		t.Error("FAILED")
	}

	if len(ht.CreateHashTable()) == ht.GetHashGroups() {
		t.Log("PASSED.")
	} else {
		t.Error("FAILED")
	}
}

func TestInsertValue(t *testing.T) {
	var ht HashTable
	hashGroups := 73
	ht.SetHashGroups(hashGroups)
	ht.CreateHashTable()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < hashGroups; i++ {
		if ht.InsertValue(rand.Intn(1000)) == false {
			t.Error("FAILED.")
		}
	}
}

func TestFindValue(t *testing.T) {
	var ht HashTable
	ht.SetHashGroups(7)
	ht.CreateHashTable()
	tests := []int{3, 7, 10, 16, 18, 11, 17}

	for _, v := range tests {
		if ht.InsertValue(v) == false {
			t.Error("FAILED")
		}
	}

	for _, v := range tests {
		if ht.FindValue(v) == false {
			t.Errorf("FAILED")
		}
	}
}
