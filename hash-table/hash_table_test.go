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
		randomValue := rand.Intn(1000)
		_, isFound := ht.FindValue(randomValue);
		if isFound == ht.InsertValue(randomValue) {
			t.Error("FAILED.")
		}
	}
}

func TestFindValue(t *testing.T) {
	var ht HashTable
	ht.SetHashGroups(7)
	ht.CreateHashTable()
	tests := []int{3, 7, 10, 16, 18, 11, 17, 18}

	for _, v := range tests {
		ht.InsertValue(v)
	}

	if count, _ := ht.FindValue(18); count != 2 {
		t.Errorf("FAILED")
	} else {
		t.Log("PASSED")
	}
}