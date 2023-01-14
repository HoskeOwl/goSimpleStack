package main

import "testing"

type gTestNode struct {
	val int
}

func BenchmarkGenericSimpleType(b *testing.B) {
	st := NewGenericStack[int]()
	for i := 0; i < b.N; i += 1 {
		st.Push(i)
		st.Pop()
	}
}

func BenchmarkGenericCustomType(b *testing.B) {
	st := NewGenericStack[gTestNode]()
	for i := 0; i < b.N; i += 1 {
		st.Push(gTestNode{i})
		st.Pop()
	}
}

func BenchmarkGenericCustomTypePointer(b *testing.B) {
	st := NewGenericStack[*gTestNode]()
	for i := 0; i < b.N; i += 1 {
		st.Push(&gTestNode{i})
		st.Pop()
	}
}
