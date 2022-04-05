package main

import (
	"fmt"
	"math/rand"
	"time"
)

type linkedList struct {
	value   int
	pointer *linkedList
}

func main() {

	rand.Seed(time.Now().UnixNano())

	list := createList(10, nil)

	// For loop for traversal
	for ; list != nil; list = list.pointer {
		fmt.Printf("%v\n", list.value)
	}
}

func createList(num int, temp *linkedList) *linkedList {
	// if list is nil, then it creates and assigns the
	// first record to the temp variable
	if temp == nil {
		temp = &linkedList{rand.Intn(100), nil}
	}

	tempList := temp

	// ensure that five records are created and added
	// to the linked list by passing the address of the new
	// linked list to the pointer of the previous record
	for i := 0; i < num; i++ {
		t := &linkedList{rand.Intn(100), nil}
		tempList.pointer = t
		tempList = tempList.pointer
	}

	return temp
}
