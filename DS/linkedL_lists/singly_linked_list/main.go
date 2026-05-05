package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkedList) insertAtFront(data int){

	if list.head == nil {
		newNode := &Node{data: data , next: nil}
		list.head = newNode
		return
	}
	newNode := &Node{data: data, next :list.head}
	list.head = newNode

}

func (list *LinkedList) insertAfterValue(afterValue, data int){
	newNode := &Node{data: data, next: nil}

	current := list.head

	for current != nil {
		if current.data == afterValue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

	fmt.Printf("Cannot insert node, after value %d doesn't exist", afterValue)
 	fmt.Println()
}


func (list *LinkedList) insertBeforeValue(beforeValue , data int){
	if list.head == nil {
		fmt.Printf("Cannot insert node, before value %d doesn't exist", beforeValue)
 		fmt.Println()
		return
	}

	if list.head.data == beforeValue {
		newNode := &Node{data: data , next: list.head}
		list.head = newNode
		return
	}

	current := list.head

	for current != nil {
		if current.next.data == beforeValue {
			newNode := &Node{data: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next
	}
	fmt.Printf("Cannot insert node, before value %d doesn't exist", beforeValue)
 	fmt.Println()
}

func (list *LinkedList) deleteFront(){
	if list.head != nil {
		list.head = list.head.next
		fmt.Printf("Head node of linked list has been deleted\n")
		return
	}
}

func (list *LinkedList) deleteLast(){
	if list.head == nil {
  		fmt.Printf("Linked list is empty\n")
		return
	}
	if list.head.next == nil {
		list.head = nil
		fmt.Printf("Last node of linked list has been deleted\n")
		return		
	}

	var current *Node = list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil

	fmt.Printf("Last node of linked list has been deleted")
}

func (list *LinkedList) deleteAfterValue(afterValue int){
	var current *Node = list.head
	

}


