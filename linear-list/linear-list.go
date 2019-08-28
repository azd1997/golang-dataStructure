package linearlist

import (
	"reflect"
	"errors"
)

// MAXSIZE 线性表的最大长度
const MAXSIZE = 10

// Elem 线性表所存储的元素应为同类型，在需要比较元素类型时需要去获取元素类型来做判断。
// 线性表元素目前仅考虑支持基础类型（Bool、String、Intx）以及聚合类型（Array、Struct）
//type Elem int
type Elem interface{}

// LinearList 线性表接口。线性表索引i从1开始
type LinearList interface {
	// 返回线性表使用长度
	Length() (len int)
	// 根据元素索引找到元素并返回元素内存位置
	FindByIndex(i int) (elemAddr *Elem, err error)
	// 根据元素值找到该元素的内存地址信息（如果只考虑顺序表，可以返回数组下表。但是我们还要考虑链式表）
	FindByValue(elemValue Elem) (elemAddrs []*Elem)
	// 插入元素到线性表第i个位置
	Insert(i int, elem Elem) (err error)
	// 删除第i个位置的元素
	Delete(i int) (err error)
	// 从前向后遍历线性表中所有元素并打印
	PrintAll()
	// 判断线性表是否为空
	IsEmpty() bool
	// 判断线性表是否已满
	IsFull() bool
	// 清空线性表所有元素，但保留线性表
	Clear() (err error)
	// 销毁线性表，释放内存空间
	Destroy() (err error)
}

// ErrNotSameType 两个元素类型不同不能用于比较！
var ErrNotSameType = errors.New("Not same type and cannot compare")
// ErrIndexOutOfRange 索引超出线性表所存数据索引范围，而不是超出底层数组索引范围
var ErrIndexOutOfRange = errors.New("This index is out of range")
// ErrNotSupportedType 不支持的元素类型
var ErrNotSupportedType = errors.New("Not supported type to compare")
// ErrNoExtraSpace 线性表已满
var ErrNoExtraSpace = errors.New("There is no extra space in the linearlist")

// A、B两元素比较大小
// const (
// 	ALarger  = 1
// 	AEqualB = 0
// 	ASmaller = -1
// )

// IsTwoElemEqual1 比较两个Elem(interface{})是否相等，true相等，false不等
func IsTwoElemEqual1(elem1, elem2 Elem) (bool, error) {
	// 比较两个Elem
	v1, v2 := reflect.ValueOf(elem1), reflect.ValueOf(elem2)
	t1, t2 := v1.Type(), v2.Type()
	k1 := v1.Kind()

	if t1 != t2 {
		return false, ErrNotSameType
	}

	// 如果两个元素类型type一致，那么分类kind一定一致
	// 之所以检查kind而不是type，是因为有的时候会使用类型别名，就像我这里使用Elem替代interface{}

	switch k1 {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return v1 == v2, nil
	case reflect.String:
		return v1 == v2, nil
	case reflect.Array, reflect.Slice:
		return v1 == v2, nil
	case reflect.Struct:
		return v1 == v2, nil
	}
	// TODO: 改正！reflect.value使用“==”比较，只是比较value内容，不会比较其潜在的数据值

	return false, nil
}

// IsTwoElemEqual2 比较两个Elem(interface{})是否相等，true相等，false不等
func IsTwoElemEqual2(elem1, elem2 Elem) bool {
	return reflect.DeepEqual(elem1, elem2)
}

// IsTwoElemEqual 比较两个Elem(interface{})是否相等，true相等，false不等
func IsTwoElemEqual(elem1, elem2 Elem) (bool, error) {
	if !IsSameType(elem1, elem2) {
		return false, ErrNotSameType
	}

	switch elem1.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return elem1 == elem2, nil
	// case string:
	// 	return elem1 == elem2, nil
	default:
		return false, ErrNotSupportedType
	}
}



// IsSameType 判断两个Elem具体类型是否一致
func IsSameType(elem1, elem2 Elem) bool {
	v1, v2 := reflect.ValueOf(elem1), reflect.ValueOf(elem2)
	t1, t2 := v1.Type(), v2.Type()

	if t1 == t2 {
		return true
	}
	return false
}