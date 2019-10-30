package ex5_array

import (
	"fmt"
	"testing"
)

var testInsert = []struct{
	array []int	// cap = 5
	index int
	value int
	result []int
}{
	{[]int{}, 0, 0, []int{}},
}

func TestStaticArray(t *testing.T) {
	s := NewStaticArray(5) // len = 0, cap = 5}
	fmt.Printf("len = %d, cap = %d\n", s.Len(), s.Cap())

	var err error
	//if err = s.Insert(2, 99); err == nil {
	//	t.Error("本应出错")
	//}

	// 正常插入五个数
	if err = s.Insert(0, 90); err != nil {
		t.Error(err)
	}
	fmt.Printf("len = %d, cap = %d\n", s.Len(), s.Cap())
	s.Print()
	if err = s.Insert(1, 91); err != nil {
		t.Error(err)
	}
	fmt.Printf("len = %d, cap = %d\n", s.Len(), s.Cap())
	s.Print()
	if err = s.Insert(2, 92); err != nil {
		t.Error(err)
	}
	fmt.Printf("len = %d, cap = %d\n", s.Len(), s.Cap())
	s.Print()
	if err = s.Insert(0, 93); err != nil {
		t.Error(err)
	}
	fmt.Printf("len = %d, cap = %d\n", s.Len(), s.Cap())
	s.Print()
	if err = s.Insert(3, 94); err != nil {
		t.Error(err)
	}
	fmt.Printf("len = %d, cap = %d\n", s.Len(), s.Cap())
	s.Print()

	// 试试满了还插
	if err = s.Insert(0, 90); err == nil {
		t.Error("本应出错")
	}
	s.Print()

	// 更改
	if err = s.Update(0, 99); err != nil {
		t.Error(err)
	}
	s.Print()

	// 删除
	if err = s.Del(3); err != nil {
		t.Error(err)
	}
	s.Print()

	return
}