package main

import "testing"

type iTestNode struct {
	val int
}

func BenchmarkInterfaceSimpleType(b *testing.B) {
	st := NewInterfaceStack()
	for i := 0; i < b.N; i += 1 {
		st.Push(i)
		if _, ok := st.Pop().(int); !ok {
			panic("Wrong type of data in stack")
		}
	}
}

func BenchmarkInterfaceSimpleCustomType(b *testing.B) {
	st := NewInterfaceStack()
	for i := 0; i < b.N; i += 1 {
		st.Push(iTestNode{i})
		if _, ok := st.Pop().(iTestNode); !ok {
			panic("Wrong type of data in stack")
		}
	}
}

func BenchmarkInterfaceCustomTypePointer(b *testing.B) {
	st := NewInterfaceStack()
	for i := 0; i < b.N; i += 1 {
		st.Push(&iTestNode{i})
		if _, ok := st.Pop().(*iTestNode); !ok {
			panic("Wrong type of data in stack")
		}
	}
}
