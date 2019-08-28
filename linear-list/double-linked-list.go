package linearlist

import (
	"log"
)

var _ LinearList = (*DoubleLinkedList)(nil)

// DoubleLinkedList 双链表
type DoubleLinkedList struct {
	head *DNode
	tail *DNode
	length int
}

// DNode 双链表节点
type DNode struct {
	elem Elem
	prior, next *DNode
}

// Length 获取线性表长度
func (l *DoubleLinkedList) Length() (len int) {
	return l.length
}

// IsEmpty 判断线性表为空
func (l *DoubleLinkedList) IsEmpty() bool {
	return l.length == 0
}

// IsFull 判断线性表是否已满
func (l *DoubleLinkedList) IsFull() bool {
	return l.length >= MAXSIZE
}

// FindByIndex 根据索引找到节点元素地址
func (l *DoubleLinkedList) FindByIndex(i int) (elemAddr *Elem, err error) {
	// 检查非空以及索引i存在
	// length >= 0。i需要在1～length之间
	if i <= 0 || i > l.length {
		return nil, ErrIndexOutOfRange
	}

	// 从head一直向后遍历. head指向i=1，那么要想按序查到第i个元素，循环i-1次即可
	// 也可以从tail向前遍历
	next := l.head
	for j:=1;j<i;j++ {
		next = next.next
	}
	// 现在next指针指向第i个节点
	return &next.elem, nil
}

// FindByValue 根据元素值查找所有符合的元素的地址
func (l *DoubleLinkedList) FindByValue(elemValue Elem) (elemAddrs []*Elem) {
	next := l.head
	for i:=1;i<=l.length;i++ {
		if next.elem == elemValue {  // 注意这里要求内部元素类型是可以使用“==”符号的
			elemAddrs = append(elemAddrs, &next.elem)
		}
		if next.next == nil {
			break
		}
		next = next.next
	}

	return elemAddrs
}

// Insert 插入元素到指定位置
func (l *DoubleLinkedList) Insert(i int, elem Elem) (err error) {

	if l.IsFull() {
		return ErrNoExtraSpace
	}
	// 注意这里是可以插入到最后位置的，所以应该是i>l.length+1才报错
	// 也就是说，l.Insert(l.Length()+1, elem)实现追加元素
	if i < 1 || i > l.length+1 {
		return ErrIndexOutOfRange
	}

	// 这里为啥部先把node构建起来？因为这三种情况下node长得不一样
	// 当然这里调用Append/Prepend是多做了索引验证的，不过这里不讲究这些细节了。
	// 当然也可以先创建DNode，不过前后指针需要先置空
	switch i {
	case 1:
		if err := l.Prepend(elem); err != nil {
			return err
		}
	case l.length+1:
		if err := l.Append(elem); err != nil {
			return err
		}
	default:  // i=2,3,...,length
		// 获取原本第i-1个元素的地址信息
		// 如果只使用FindByIndex呢？可以利用unsafe.SizeOf(var)来推断出next指针位置（因为是相邻存储的）
		frontNode, err := l.FindNodeByIndex(i-1)
		if err != nil {
			return err
		}

		// 插入第i个节点
		// 注意这里如果是插入第length+1个节点，next指针会是空，这不影响，但是需要更新tail，同样的如果是第一个节点，也需要更新head
		node := DNode{
			elem: elem,
			next: frontNode.next,
			prior: frontNode,
		}

		// 这里有个特殊情况：当length=1时，要插入

		// 修改原第i个节点，再修改原第i-1个节点
		frontNode.next.prior = &node
		frontNode.next = &node

		// length + 1
		l.length++
	}

	return nil
}

// Delete 删除第i个位置的元素
// 最好时间复杂度： 删除最后一个位置，不用挪动数据，O(1)
// 最坏时间复杂度： 删除第一个位置，全挪，O(n)
// 平均时间复杂度： (1+2+...+n)/n = (n+1)/2， O(n)
func (l *DoubleLinkedList) Delete(i int) (err error) {
	// 检查所要删除的索引是存在的
	if i < 1 || i > l.length {
		return ErrIndexOutOfRange
	}

	// 删除第i个节点可以有两种方式：
	// 1.找到第i-1个节点，从第i-1个节点连到第i+1个节点
	// 2.找到第i个节点，将第i+1个节点的内容覆盖掉第i个节点

	// 借助临时节点指针实现
	switch i {
	case 1:
		// 回收原来的第一个节点空间
		// 更新head，还要更新原第2个节点。这里先后顺序不重要，相应调整就好
		l.head = l.head.next
		l.head.prior = nil
	case l.length:
		// 更新tail， 并回收最后一个节点
		// 双链表在这里有个好处是不需要再去Find查找
		l.tail = l.tail.prior
		l.tail.next = nil
	default:
		frontNode, err := l.FindNodeByIndex(i-1)
		if err != nil {
			return err
		}
		// 回收第i个
		// 更新第i-1个的next指向， 更新原第i+1个节点的prior
		frontNode.next = frontNode.next.next
		frontNode.next.prior = frontNode
	}

	// length - 1
	l.length--

	return nil
}

// PrintAll 从前向后遍历线性表中所有元素并打印
func (l *DoubleLinkedList) PrintAll() {
	// 检查l长度，至少要有一个元素
	if l.length == 0 {
		log.Println("双链表为空")
	}

	log.Println("双链表元素：")
	// 迭代
	// 这里由于是双链表，其实可以开两个协程去做，从两头往中间遍历。这里就不这么折腾了。
	next := l.head
	for j:=0;j<l.length;j++ {
		log.Printf("%v\r", next.elem)
		// 这里需要先检查下next是否为空，不然的话next=next.next会报错
		if next.next == nil {
			break
		}
		next = next.next
	}
	log.Println()
}

// Clear 清除线性表所有元素，但不回收内存
func (l *DoubleLinkedList) Clear() (err error) {

	return nil
}

// Destroy 销毁线性表，回收内存
func (l *DoubleLinkedList) Destroy() (err error) {

	return nil
}

// Append 末尾追加元素
func (l *DoubleLinkedList) Append(elem Elem) (err error) {

	// 检查单链表是否已满
	// 这里不像Insert，可以用IsFull来检查有没有空间
	if l.IsFull() {
		return ErrNoExtraSpace
	}

	// 构建节点
	node := DNode{
		elem: elem,
		next: nil,
		prior: l.tail,
	}

	// 更新原先最后一个节点的next
	l.tail.next = &node
	// 更新tail
	l.tail = &node
	// 更新length
	l.length++

	return nil
}

// Prepend 头部追加元素
func (l *DoubleLinkedList) Prepend(elem Elem) (err error) {

	// 检查单链表是否已满
	if l.IsFull() {
		return ErrNoExtraSpace
	}

	// 构建节点
	node := DNode{
		elem: elem,
		next: l.head,	// 指向原先的第一个节点
		prior: nil,
	}

	// Prepend需要考虑一个特殊情况：插入第一个节点时，其既是head也是tail
	if l.length == 0 {
		// 这种情况下原l.head为nil，所以不影响tail直接指向node
		l.head = &node
		l.tail = &node
	} else {
		// 更新原第1个节点再更新head
		l.head.prior = &node
		l.head = &node
	}

	// 更新length
	l.length++

	return nil
}

// Set方法基于Find方法实现

// FindNodeByIndex 根据索引返回节点指针
func (l *DoubleLinkedList) FindNodeByIndex(i int) (*DNode, error) {
	// 检查非空以及索引i存在
	// length >= 0。i需要在1～length之间
	if i <= 0 || i > l.length {
		return nil, ErrIndexOutOfRange
	}

	// 根据查找的位置更靠近头部还是更靠近尾部决定从哪头开始遍历
	if 2*i <= l.length {
		// 更靠近头部，从头部开始遍历
		// 从head一直向后遍历. head指向i=1，那么要想按序查到第i个元素，循环i-1次即可
		next := l.head
		for j:=1;j<i;j++ {
			next = next.next
		}
		// 现在next指针指向第i个节点
		return next, nil
	}
	// 否则更靠近尾部，从尾部向前遍历
	prior := l.tail
	for j:=l.length;j>i;j-- {
		prior = prior.prior
	}
	return prior, nil
}