package linearlist

// MAXSIZE 线性表的最大长度
const MAXSIZE = 20

// Elem 线性表所存储的元素应为同类型，在此处限定其为int。
type Elem int

// LinearList 线性表接口
type LinearList interface {
	// 初始化线性表，构建空线性表并插入一个元素
	Init(i int, elem Elem)
	// 返回线性表使用长度
	Length() (len int)
	// 根据元素索引找到元素并返回元素内存位置
	FindByIndex(i int) (elemAddr *Elem, err error)
	// 根据元素值找到该元素的内存地址信息（如果只考虑顺序表，可以返回数组下表。但是我们还要考虑链式表）
	FindByValue(elemValue Elem) (elemAddr *Elem, err error)
	// 插入元素到线性表第i个位置
	Insert(i int, elem Elem) (newLen int, err error)
	// 删除第i个位置的元素
	Delete(i int) (newLen int, err error)
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