package main

type (
	gNode[NT any] struct {
		val  NT
		next *gNode[NT]
	}
	GenericStack[ST any] struct {
		top *gNode[ST]
		len int
	}
)

func NewGenericStack[GS any]() *GenericStack[GS] {
	return &GenericStack[GS]{}
}

func (gstack *GenericStack[ST]) Push(val ST) {
	var n gNode[ST] = gNode[ST]{val: val, next: gstack.top}
	gstack.len += 1
	gstack.top = &n
}

func (gstack *GenericStack[ST]) Pop() (res ST, exists bool) {
	if gstack.len <= 0 {
		exists = false
		return
	}
	gstack.len -= 1

	var n *gNode[ST] = gstack.top
	gstack.top = n.next
	return n.val, true
}

func (gstack *GenericStack[ST]) Peak() (res ST, exists bool) {
	if gstack.len <= 0 {
		exists = false
		return
	}
	return gstack.top.val, true
}

func (gstack GenericStack[ST]) Len() int {
	return gstack.len
}
