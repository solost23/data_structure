package set

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet(1e7 + 20)

	type arg struct {
		Name string
	}
	set.Add(1, strconv.Itoa(2), arg{Name: "alex"})
	// 查看1是否存在
	if !set.Has(1) {
		t.Errorf("1 exitst")
	}
	// 查看容量
	if set.Len() != 3 {
		t.Errorf("set len is 3")
	}
	// 删除1
	set.Remove(1)
	if set.Len() != 2 {
		t.Errorf("set len is 2")
	}

	set.Clear()
	set.Add(1)
	if set.Len() != 1 {
		t.Errorf("set len is 1")
	}
}

func TestSetOp(t *testing.T) {
	s1 := NewSet(10)
	s1.Add(1, 2, 3)
	s2 := NewSet(10)
	s2.Add(3, 4, 5, )

	result := s1.Union(s2)
	fmt.Println(result.set)

	result1 := s1.Diff(s2)
	fmt.Println(result1.set)

	result3 := s1.Mixed(s2)
	fmt.Println(result3.set)
}

func BenchmarkSet(B *testing.B) {
	set := NewSet(1e5)
	for i := 0; i < B.N; i++ {
		args := []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		set.Add(args...)
		for i := 0; i < len(args); i++ {
			set.Remove(args[i])
		}
	}
}
