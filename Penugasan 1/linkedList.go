package main

import (
	"errors"
	"fmt"
)

type Linked struct {
	Data int
	Next *Linked
}

type List struct {
	Head *Linked
	Tail *Linked
	size uint
}

func (l *List) PushBack(data int) {
	node := &Linked{
		Data: data,
		Next: nil,
	}

	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}

	l.size++
}

func (l *List) PopBack() error {
	if l.Head == nil {
		return errors.New("error gan")
	}

	if l.Head.Data == l.Tail.Data {
		l.Head = nil
		l.Tail = nil
		return nil
	} else {
		current := l.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		l.Tail = current
		current.Next = nil
	}
	l.size--
	return nil
}

func (l *List) DisplayAll() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	SingleList := List{}
	var choice int
	var data int

	for {
		fmt.Println("\n1. Push Back")
		fmt.Println("2. Pop Back")
		fmt.Println("3. Display All")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter value: ")
			fmt.Scanln(&data)
			SingleList.PushBack(data)
			SingleList.DisplayAll()
		case 2:
			SingleList.PopBack()
			SingleList.DisplayAll()
		case 3:
			SingleList.DisplayAll()
		case 4:
			return
		default:
			fmt.Println("1-3 aja we")
		}
	}
}
