package main

import (
	"errors"
	"fmt"
)

type HashItem struct {
	key   string
	value int
}

type HashTable struct {
	size  int
	slots []HashItem
	count int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		slots: make([]HashItem, size),
		count: 0,
	}
}

func (h *HashTable) hash(key string) int {
	mult := 1
	hv := 0
	for _, c := range key {
		hv += mult * int(c)
		mult++
	}

	return hv % h.size
}

func (h *HashTable) PutLinear(key string, value int) {
	hv := h.hash(key)
	for h.slots[hv].key != "" && h.slots[hv].key != key {
		hv = (hv + 1) % h.size
	}

	if h.slots[hv].key == "" {
		h.count++
	}

	h.slots[hv] = HashItem{key: key, value: value}
}

func (h *HashTable) Get(key string) (int, error) {
	hv := h.hash(key)
	for h.slots[hv].key != "" && h.slots[hv].key != key {
		hv = (hv + 1) % h.size
	}

	if h.slots[hv].key == "" {
		return 0, errors.New("key not registred")
	} else {
		return h.slots[hv].value, nil
	}
}

func main() {
	table := NewHashTable(100)
	fmt.Println("inserting item")
	table.PutLinear("hello world", 10)
	fmt.Println("item inserted")
	fmt.Println("getting Item")
	value, err := table.Get("hello world")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}
