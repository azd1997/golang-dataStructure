package linearlist

// SequenceList 顺序表 顺序存储的线性表
type SequenceList struct {
	elems [MAXSIZE]Elem
	length int
}

// Init 初始化一个空的顺序表
func (l *SequenceList) Init(i int, elem Elem) {

}

// Length 返回线性表使用长度
func (l *SequenceList) Length() (len int) {
	return l.length
}

// FindByIndex 根据元素索引找到元素并返回元素内存位置
func (l *SequenceList) FindByIndex(i int) (elemAddr *Elem, err error) {


	return elemAddr, nil
}

// FindByValue 根据元素值找到该元素的内存地址信息（如果只考虑顺序表，可以返回数组下表。但是我们还要考虑链式表）
func (l *SequenceList) FindByValue(elemValue Elem) (elemAddr *Elem, err error) {


	return elemAddr, nil
}

// Insert 插入元素到线性表第i个位置
func (l *SequenceList) Insert(i int, elem Elem) (err error) {

	return nil
}

// Delete 删除第i个位置的元素
func (l *SequenceList) Delete(i int) (err error) {

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

// PrintAll 从前向后遍历线性表中所有元素并打印
func (l *SequenceList) PrintAll() {

}

// IsEmpty 判断线性表是否为空
func (l *SequenceList) IsEmpty() bool {

	return false
}

// IsFull 判断线性表是否已满
func (l *SequenceList) IsFull() bool {

	return false
}

// Clear 清除线性表所有元素，但不回收内存
func (l *SequenceList) Clear() (err error) {

	return nil
}

// Destroy 销毁线性表，回收内存
func (l *SequenceList) Destroy() (err error) {

	return nil
}

