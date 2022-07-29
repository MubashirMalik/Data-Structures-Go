package hashtable

import (
	ll "github.com/MubashirMalik/Data-Structures-Go/linked-list"
)

var hashGroups int
type HashTable []ll.LinkedList
var hashTable HashTable
var valueCount HashTable

func (ht *HashTable) SetHashGroups(hg int) bool {
	if hashGroups == 0 && hg > 0 {
		hashGroups = hg
		return true
	}
	return false
}

func (ht *HashTable) GetHashGroups() int {
	return hashGroups
}

func (ht *HashTable) CreateHashTable() []ll.LinkedList {
	hashTable = make([]ll.LinkedList, hashGroups)
	valueCount = make([]ll.LinkedList, hashGroups)
	return hashTable
}

func (ht *HashTable) InsertValue(value int) bool {
	idx := hashTable[value % hashGroups].Len()
	// Don't insert if already present in hash map
	if idxFound, isFound := hashTable[value % hashGroups].Find(value); isFound {
		count, _ := valueCount[value % hashGroups].DataAt(idxFound)
		valueCount[value % hashGroups].InsertAt(count+1, idxFound) 
		return false;
	}
	valueCount[value % hashGroups].InsertAt(1, idx)
	return hashTable[value % hashGroups].InsertAt(value, idx)
}

func (ht *HashTable) GetHashTable() []ll.LinkedList {
	return hashTable
}

func (ht *HashTable) FindValue(value int) (int, bool) {
	idx, isFound := hashTable[value % hashGroups].Find(value)
	count, _ := valueCount[value % hashGroups].DataAt(idx)
	return count, isFound
}

func (ht *HashTable) RemoveValue(value int) (int, bool) {
	idx, isFound := hashTable[value % hashGroups].Find(value)
	if (!isFound) {
		return 0, false
	}
	count, _ := valueCount[value % hashGroups].DataAt(idx)
	if (count == 1) {
		valueCount[value % hashGroups].DeleteAt(idx)
		return hashTable[value % hashGroups].DeleteAt(idx)
	} else {
		valueCount[value % hashGroups].InsertAt(count-1, idx)
		return count, true
	}
}