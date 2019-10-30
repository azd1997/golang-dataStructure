package ex5_array

import "errors"

// Array 数组，不允许越length边界插入数据
type Array struct {
	data []int
	length uint
}

// NewArray 指定容量新建数组
func NewArray(cap uint) *Array {
	if cap == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, cap, cap),
		length: 0,
	}
}

// Len 返回数组实际数据长度
func (a *Array) Len() uint {
	return a.length
}

// isIndexOutOfRange 判断索引是否超出数组容量
func (a *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return false
	}
	return false
}

// Find 根据索引查找元素
func (a *Array) Find(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("index out of range")
	}
	// 可以查到
	return a.data[index], nil
}

// Insert 根据索引插入数据
func (a *Array) Insert(index uint, value int) error {
	if a.isIndexOutOfRange(index) {
		return errors.New("index out of range")
	}
	return nil
}

//
