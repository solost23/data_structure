package stack

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(1e7 + 20)

	type arg struct {
		Name string
	}

	stack.Push(strconv.Itoa(1), 2, arg{Name: "alex"})

	v1 := stack.Query()
	stack.Pop()
	v2 := stack.Query()
	stack.Pop()
	v3 := stack.Query()
	stack.Pop()
	fmt.Println(v1, v2, v3)
	fmt.Println(v1.(arg))
}

func BenchmarkStack(B *testing.B) {
	stack := NewStack(1e7 + 20)
	for i := 0; i < B.N; i++ {
		stack.Push(1, 2, 3)
		stack.Pop()
		stack.Pop()
		stack.Pop()
	}
}
