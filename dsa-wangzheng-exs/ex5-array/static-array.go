package ex5_array

import (
	"errors"
	"fmt"
)

// StaticArray 使用固定容量的切片来模拟固定长度的数组
type StaticArray struct {
	data []int
}

// NewStaticArray 根据容量生成固定容量的数组
func NewStaticArray(cap uint) *StaticArray {
	if cap == 0 {
		return nil
	}
	return &StaticArray{make([]int, 0, cap)}
}

// Cap 数组容量
func (a *StaticArray) Cap() int {
	return cap(a.data)
}

// Len 数组已使用长度
func (a *StaticArray) Len() int {
	return len(a.data)
}

// Find 根据索引查找元素，支持负数索引
func (a *StaticArray) Find(index int) (res int, err error) {
	l := a.Len()
	err = isIndexOut(index, l, OP_FIND)
	if err != nil {
		return 0, err
	}
	var realIndex int
	realIndex = toRealIndex(index, l)

	return a.data[realIndex], nil
}

// Insert 根据下标插入元素，保持原元素的相对位置。 有一种取巧的插入是把直接把插入位置上的元素搬到末尾，这里不用
func (a *StaticArray) Insert(index int, elem int) (err error) {
	// 检查数组是否已满
	if a.isFull() {
		return errors.New("array is full")
	}

	// 插入到数组所使用部分的末尾。append
	l := a.Len()
	if index == l || index == -1 {
		a.data = append(a.data, elem)
		return nil
	}

	// 一般位置插入(检查越界，转换索引，插入)
	err = isIndexOut(index, l, OP_INSERT)
	if err != nil {
		return err
	}
	var realIndex int
	realIndex = toRealIndex(index, l)

	// 在真正的插入之前  记得要给数组追加一个元素，用来给插入占位置，不然插入会索引越界
	// 并且每次报错时，应该重新将这个位置删除。（这个设计太操蛋了，还是加个length好）
	a.data = append(a.data, 0)

	// 原数据向后挪移
	for i:=l-1; i>=realIndex; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[realIndex] = elem

	return nil
}

// Del 删数据，后续前移
// 注意：在这样的没有length标记字段的设计里，没办法回收空间，删除只能通过重新分配内存实现
func (a *StaticArray) Del(index int) error {
	l := a.Len()
	var err error
	err = isIndexOut(index, l, OP_DELETE)
	if err != nil {
		return err
	}
	var realIndex int
	realIndex = toRealIndex(index, l)

	// 删除通过拷贝另一个切片实现
	slice := make([]int, l-1, a.Cap())
	for i:=0; i<realIndex; i++ {
		slice[i] = a.data[i]
	}
	for i:=realIndex; i<l-1; i++ {
		slice[i] = a.data[i+1]
	}
	a.data = slice
	return nil
}

// Update 更新元素值
func (a *StaticArray) Update(index int, new int) error {
	l := a.Len()
	var err error
	err = isIndexOut(index, l, OP_DELETE)
	if err != nil {
		return err
	}
	var realIndex int
	realIndex = toRealIndex(index, l)

	a.data[realIndex] = new
	return nil
}

// Print 打印
func (a *StaticArray) Print() {
	fmt.Println(a.data)
}

const (
	OP_INSERT = iota
	OP_DELETE
	OP_UPDATE
	OP_FIND
)

var opMap = map[uint8]string{
	OP_INSERT : "op_insert",
	OP_DELETE : "op_delete",
	OP_UPDATE : "op_update",
	OP_FIND : "op_find",
}

// isIndexOut 判断索引是否越界
func isIndexOut(index int, listLength int, op uint8) (err error) {

	switch op {
	case OP_INSERT:
		if index > listLength || index < -listLength-1 {
			err = errors.New("index out of range")
		}
	case OP_DELETE, OP_UPDATE, OP_FIND:
		if index >= listLength || index < -listLength {
			err = errors.New("index out of range")
		}
	default:
		err = errors.New("unknown operation")
	}

	if err != nil {
		return fmt.Errorf("%s: %s", opMap[op], err)
	}
	return nil
}

// toRealIndex 处理int型索引下标，将其转换为自然数。调用之前需根据不同操作检查索引越界
func toRealIndex(index int, listLength int) (realIndex int) {

	// 0, 1, 2, 3 中间插入4 => 0, 1, 4, 2, 3
	// 现长度4， 插入正索引2， 插入负索引-3
	// 0, 1, 2, 3 中间插入4 => 0, 1, 4, 2, 3
	// 现长度4， 插入正索引4， 插入负索引-1
	if index < 0 {
		realIndex = listLength + index + 1
	} else {
		realIndex = index
	}
	return realIndex
}

// isFull 判断数组是否已满
func (a *StaticArray) isFull() bool {
	return a.Len() == a.Cap()
}