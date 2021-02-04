package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func (list *linkedList) append(data int) {
	//creates a new node that will be added to the linkedList
	newNode := &node{
		data: data,
	}

	//checks if the head of the list is
	//empty if so it adds the node to the head and tail
	//if not it links the tail next to the new node and changes the tail to the new node
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}

	//increase the length value of the linked list by one
	list.length++
}

func (list *linkedList) remove(data int) {
	//If it is the head, remove it
	if list.head.data == data {
		list.head = list.head.next
		list.length--
		return
	}

	currentNode := list.head

	for currentNode.next != nil {
		//If the next node is equal to the data and it is the tail remove
		if currentNode.next == list.tail && list.tail.data == data {
			list.tail = currentNode
			list.length--
			break
		}

		//If the next node is equal to the data remove
		if currentNode.next.data == data {
			currentNode.next = currentNode.next.next
			list.length--
			break
		}

		// currentNode should be the next
		currentNode = currentNode.next
	}
}

func (list *linkedList) contains(data int) bool {

	if list.head.data == data || list.tail.data == data {
		return true
	}

	currentNode := list.head

	for currentNode.next != nil {
		if currentNode.data == data {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

func (list *linkedList) get(data int) *node {
	if list.head.data == data {
		return list.head
	}

	if list.tail.data == data {
		return list.tail
	}

	currentNode := list.head

	for currentNode.next != nil {
		if currentNode.data == data {
			return currentNode
		}

		currentNode = currentNode.next
	}

	return nil
}

func (list *linkedList) printAll() {
	nextPrint := list.head

	fmt.Println("--------------------")
	for i := list.length; i > 0; i-- {
		fmt.Printf("%d \n", nextPrint.data)
		nextPrint = nextPrint.next
	}

	fmt.Println("--------------------")

}

func main() {

	ll := &linkedList{}
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)
	ll.append(6)

	fmt.Printf("length is  %d \n", ll.length) //length is 6
	fmt.Printf("head is %d \n", ll.head.data) //head is 1
	fmt.Printf("tail is %d \n", ll.tail.data) //tail is 6

	ll.remove(2)

	fmt.Printf("contains 2 = %t \n", ll.contains(2)) //contains 2 = false
	fmt.Printf("contains 3 = %t \n", ll.contains(3)) //contains 3 = true

	n := ll.get(1)
	if n != nil {
		fmt.Printf("Found Node %d \n", n.data) //Found Node 1
	}

	n = ll.get(12)
	if n == nil {
		fmt.Printf("No node found containing %d \n", 12) //No node found containing 12
	}

	ll.printAll()
}
