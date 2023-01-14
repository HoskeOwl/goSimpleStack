package main

import "fmt"

func main() {
	istack := NewInterfaceStack()
	istack.Push(12)
	istack.Push(32)
	if val, ok := istack.Pop().(int); !ok {
		panic("wrong type in interface stack")
	} else {
		fmt.Printf("Got '%v' from interface stack\n", val)
	}

	gstack := NewGenericStack[int]()
	gstack.Push(54)
	gstack.Push(67)
	if val, exists := gstack.Pop(); exists {
		fmt.Printf("Got '%v' from generic stack\n", val)
	}
}
