package main

type (
	iNode struct {
		val  interface{}
		next *iNode
	}
	InterfaceStack struct {
		top *iNode
		len int
	}
)

func NewInterfaceStack() *InterfaceStack {
	return &InterfaceStack{}
}

func (istack *InterfaceStack) Push(val interface{}) {
	var n iNode = iNode{val: val, next: istack.top}
	istack.len += 1
	istack.top = &n
}

func (istack *InterfaceStack) Pop() interface{} {
	if istack.len <= 0 {
		return nil
	}
	istack.len -= 1
	var n *iNode = istack.top
	istack.top = n.next
	return n.val
}

func (istack *InterfaceStack) Peak() interface{} {
	if istack.len <= 0 {
		return nil
	}
	return istack.top.val
}

func (istack *InterfaceStack) Len() int {
	return istack.len
}
