package hashtable

import (
	ll "github.com/MubashirMalik/Data-Structures-Go/linked-list"
)

var hashGroups int
type HashTable []ll.LinkedList
var hashTable HashTable

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
	return hashTable
}

func (ht *HashTable) InsertValue(value int) bool {
	idx := hashTable[value % hashGroups].Len()
	// Don't insert if already present in hash map
	if _, isFound := hashTable[value % hashGroups].Find(value); isFound {
		return false;
	}
	return hashTable[value % hashGroups].InsertAt(value, idx)
}

func (ht *HashTable) GetHashTable() []ll.LinkedList {
	return hashTable
}

func (ht *HashTable) FindValue(value int) bool {
	_, isFound := hashTable[value % hashGroups].Find(value)
	return isFound
}
