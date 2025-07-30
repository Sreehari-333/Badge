package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	Head *Node
}

// Inserting elements to Linked List

func (list *LinkedList) insert(data int) {
	newNode := Node{data: data}

	if list.Head == nil {
		list.Head = &newNode
		return
	}

	current := list.Head

	for current.next != nil {
		current = current.next
	}
	current.next = &newNode
}

// Displaying elements in Linked List

func (list *LinkedList) display() {

	if list.Head == nil {
		fmt.Println("List is empty")
		return
	}

	current := list.Head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// Searching

func (list *LinkedList) search(val int) bool {

	if list.Head == nil {
		fmt.Println("List is empty")
		return false
	}

	current := list.Head

	for current != nil {
		if current.data == val {
			fmt.Println("Value is in the list")
			return true
		}
		current = current.next
	}

	return false
}

func main() {
	l1 := &LinkedList{}

	l1.insert(10)
	l1.insert(20)
	l1.insert(30)
	l1.insert(40)

	l1.display()
	l1.search(40)

}
