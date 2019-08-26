package linearlist

import (
	"log"
	"errors"
)

// SequenceList 顺序表 顺序存储的线性表
// l := SequenceList{} l: {[<nil> <nil ... >], 0}
// 要想实现动态扩容，底层基于slice，或者自己写array的自动扩容。
type SequenceList struct {
	// elems为数组
	elems [MAXSIZE]Elem
	// 长度需要保持与elems长度一致，之所以这么写，主要是为了照顾链表
	length int
}

// Init 初始化一个空的顺序表
// l := &SequenceList{}
// TODO: DEPRECATED!!!
func (l *SequenceList) Init(i int, elem Elem) {

}

// Length 返回线性表使用长度
func (l *SequenceList) Length() (len int) {
	return l.length
}

// FindByIndex 根据元素索引找到元素并返回元素内存位置
func (l *SequenceList) FindByIndex(i int) (elemAddr *Elem, err error) {
	// 检查非空以及索引i存在
	// length >= 0。i需要在1～length之间
	if i <= 0 || i > l.length {
		return nil, errors.New("该索引不存在！")
	}
	elemAddr = &l.elems[i-1]
	return elemAddr, nil
}

// FindByValue 根据元素值找到该元素的内存地址信息（如果只考虑顺序表，可以返回数组下表。但是我们还要考虑链式表）
func (l *SequenceList) FindByValue(elemValue Elem) (elemAddrs []*Elem, err error) {
	// 循环遍历线性表， 匹配值
	for i:=1;i<=l.length;i++ {
		if l.elems[i-1] == elemValue {  // 注意这里要求内部元素类型是可以使用“==”符号的
			elemAddrs = append(elemAddrs, &l.elems[i-1])
		}
	}

	return elemAddrs, nil
}

// Insert 插入元素到线性表第i个位置，i从1开始
func (l *SequenceList) Insert(i int, elem Elem) (err error) {
	// 要求原线性表未满、i在线性表索引值范围
	if l.IsFull() {
		return ErrNoExtraSpace
	}
	// 注意这里是可以插入到最后位置的，所以应该是i>l.length+1才报错
	// 也就是说，l.Insert(l.Length()+1, elem)实现追加元素
	if i < 1 || i > l.length+1 {
		return ErrIndexOutOfRange
	}

	// 原本第i个元素后移。注意倒序遍历。
	// 如果i=l.length+1，这段代码会被跳过
	for j:=l.length;j>=i;j-- {
		l.elems[j] = l.elems[j-1]
	}

	// 插入第i个元素
	l.elems[i-1] = elem

	// length + 1
	l.length++

	return nil
}

// Delete 删除第i个位置的元素
func (l *SequenceList) Delete(i int) (err error) {
	// 检查所要删除的索引是存在的
	if i < 1 || i > l.length {
		return ErrIndexOutOfRange
	}

	// 第i位往后所有元素前移一位，就实现删除了
	for j:=i;j<l.length;j++ {
		l.elems[j-1] = l.elems[j]
	}

	// length - 1
	l.length--

	return nil
}

// PrintAll 从前向后遍历线性表中所有元素并打印
func (l *SequenceList) PrintAll() {
	// 检查l长度，至少要有一个元素
	if l.length == 0 {
		log.Println("线性表为空")
	}

	log.Println("线性表元素：")
	for j:=0;j<l.length;j++ {
		log.Printf("%v\r", l.elems[j])
	}
	log.Println()
}

// IsEmpty 判断线性表是否为空
func (l *SequenceList) IsEmpty() bool {
	return l.length == 0
}

// IsFull 判断线性表是否已满
func (l *SequenceList) IsFull() bool {
	return l.length == MAXSIZE
}

// Clear 清除线性表所有元素，但不回收内存
func (l *SequenceList) Clear() (err error) {

	return nil
}

// Destroy 销毁线性表，回收内存
func (l *SequenceList) Destroy() (err error) {

	return nil
}

// Append 末尾追加元素
func (l *SequenceList) Append(elem Elem) (err error) {


	return nil
}

// Prepend 末尾追加元素
func (l *SequenceList) Prepend(elem Elem) (err error) {

	return nil
}

// Set方法基于Find方法实现