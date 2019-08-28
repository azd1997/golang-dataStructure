package linearlist

import (
	"log"
)

var _ LinearList = (*CircularSingleLinkedList)(nil)

// CircularSingleLinkedList 循环单链表
type CircularSingleLinkedList struct {
	// 循环单链表只需要一个head。当然也可以加一个tail，在有些时候会方便许多
	head *CSNode
	length int
}

// CSNode 循环单链表的节点
type CSNode struct {
	elem Elem
	next *CSNode
}

// Length 返回线性表使用长度
func (l *CircularSingleLinkedList) Length() (len int) {
	return l.length
}

// FindByIndex 根据元素索引找到元素并返回元素内存位置
// 假设MAXSIZE可以很大
// 最好、最坏、平均时间复杂度： O(1)
// 最好时间复杂度： 查找第一个节点元素，O(1)
// 最坏时间复杂度： 查找最后一个节点元素，O(n)
// 平均时间复杂度： (0+1+2+...+n)/n = (n+1)/2， O(n)
func (l *CircularSingleLinkedList) FindByIndex(i int) (elemAddr *Elem, err error) {
	// 检查非空以及索引i存在
	// length >= 0。i需要在1～length之间
	if i <= 0 || i > l.length {
		return nil, ErrIndexOutOfRange
	}

	// 从head一直向后遍历. head指向i=1，那么要想按序查到第i个元素，循环i-1次即可
	next := l.head
	for j:=1;j<i;j++ {
		next = next.next
	}
	// 现在next指针指向第i个节点
	return &next.elem, nil
}

// FindByValue 根据元素值找到该元素的内存地址信息（如果只考虑顺序表，可以返回数组下表。但是我们还要考虑链式表）
// 假设MAXSIZE可以很大
// 最好、最坏、平均时间复杂度： O(n)
func (l *CircularSingleLinkedList) FindByValue(elemValue Elem) (elemAddrs []*Elem) {

	// 循环遍历线性表， 匹配值
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

// Insert 插入元素到线性表第i个位置，i从1开始
// 假设MAXSIZE可以很大
// 最好时间复杂度： 插入到末尾后一个位置和第一个位置 O(1)
// 最坏时间复杂度： 插入到末尾位置，需要遍历到最后一个，O(n)
// 平均时间复杂度： (1+1+2+...+n)/(n+1) ～ n/2， O(n)
func (l *CircularSingleLinkedList) Insert(i int, elem Elem) (err error) {
	// 要求原线性表未满、i在线性表索引值范围
	// 这里需要注意IsFull： i >= l.length。但是这里和sequencelist.go的区别在于
	// 顺序表底层基于数组，索引其实是0开始，所以它可以l.IsFull就直接报错
	// 但这里的问题是索引是从1开始的，最大值和长度值一致，所以这种情况下最好不用IsFull，改为i>MAXSIZE

	// 选择IsFull还是i>Maxsize的区别在于IsFull如果满了，就绝对不应该允许插入，从效果上等价于i>=Maxsize
	if l.IsFull() {
		return ErrNoExtraSpace
	}
	// 注意这里是可以插入到最后位置的，所以应该是i>l.length+1才报错
	// 也就是说，l.Insert(l.Length()+1, elem)实现追加元素
	if i < 1 || i > l.length+1 {
		return ErrIndexOutOfRange
	}

	switch i {
		// 注意：prepend、append内部已经对length进行加一了
	case 1:
		if err := l.Prepend(elem); err != nil {
			return err
		}
	case l.length+1:
		if err := l.Append(elem); err != nil {
			return err
		}
	default:
		// 获取原本第i-1个元素的地址信息
		// 如果只使用FindByIndex呢？可以利用unsafe.SizeOf(var)来推断出下一个节点指针位置（因为是相邻存储的）
		frontNode, err := l.FindNodeByIndex(i-1)
		if err != nil {
			return err
		}

		// 插入第i个节点
		// 注意这里如果是插入第length+1个节点，next指针会是空，这不影响，但是需要更新tail，同样的如果是第一个节点，也需要更新head
		node := CSNode{
			elem: elem,
			next: frontNode.next,
		}

		// 修改frontNode
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
func (l *CircularSingleLinkedList) Delete(i int) (err error) {
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
		// 更新head
		l.head = l.head.next
	case l.length:
		// 更新tail， 并回收最后一个节点
		lastSecond, err := l.FindNodeByIndex(l.length-1)
		if err != nil {
			return err
		}
		lastSecond.next = l.head
	default:
		frontNode, err := l.FindNodeByIndex(i-1)
		if err != nil {
			return err
		}
		// 回收第i个
		// 更新第i-1个的next指向
		frontNode.next = frontNode.next.next
	}

	// length - 1
	l.length--

	return nil
}

// PrintAll 从前向后遍历线性表中所有元素并打印
func (l *CircularSingleLinkedList) PrintAll() {
	// 检查l长度，至少要有一个元素
	if l.length == 0 {
		log.Println("循环单链表为空")
	}

	log.Println("循环单链表元素：")
	//迭代
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

// IsEmpty 判断线性表是否为空
func (l *CircularSingleLinkedList) IsEmpty() bool {
	return l.length == 0
}

// IsFull 判断线性表是否已满
func (l *CircularSingleLinkedList) IsFull() bool {
	return l.length >= MAXSIZE
}

// Clear 清除线性表所有元素，但不回收内存
func (l *CircularSingleLinkedList) Clear() (err error) {

	return nil
}

// Destroy 销毁线性表，回收内存
func (l *CircularSingleLinkedList) Destroy() (err error) {

	return nil
}

// Append 末尾追加元素
func (l *CircularSingleLinkedList) Append(elem Elem) (err error) {

	// 检查单链表是否已满
	// 这里不像Insert，可以用IsFull来检查有没有空间
	if l.IsFull() {
		return ErrNoExtraSpace
	}

	// 构建节点
	node := CSNode{
		elem: elem,
		next: l.head,
	}

	// 更新原先最后一个节点的next
	originLast, err := l.FindNodeByIndex(l.length)
	if err != nil {
		return err
	}
	originLast.next = &node
	// 更新length
	l.length++

	return nil
}

// Prepend 头部追加元素
func (l *CircularSingleLinkedList) Prepend(elem Elem) (err error) {

	// 检查单链表是否已满
	if l.IsFull() {
		return ErrNoExtraSpace
	}

	// 构建节点
	node := CSNode{
		elem: elem,
		next: l.head,	// 指向原先的第一个节点
	}

	// Prepend有个特殊情况就是：插入第一个时也是调用prepend，但是此时第一个节点就是最后一个节点，最后一个节点的next需指向头结点。在我们这里，头结点其实就是&l.head
	// 但是我们这里由于没有将l与普通节点作相同结构体定义，所以不能这么写。
	// 这里做个小变动，将循环限定在数据节点之间，也就是总共有头指针（头结点）、数据节点（node1,...,noden），只在数据节点中达成循环
	// 这么做的话，当只有一个节点时，其next域指向自身，由于这里head还未更新，所以写成node.next=&node
	if l.length == 0 {
		node.next = &node
	}

	// 更新head
	l.head = &node
	// 更新length
	l.length++

	return nil
}

// Set方法基于Find方法实现

// FindNodeByIndex 根据索引返回节点指针
func (l *CircularSingleLinkedList) FindNodeByIndex(i int) (*CSNode, error) {
		// 检查非空以及索引i存在
	// length >= 0。i需要在1～length之间
	if i <= 0 || i > l.length {
		return nil, ErrIndexOutOfRange
	}

	// 从head一直向后遍历. head指向i=1，那么要想按序查到第i个元素，循环i-1次即可
	next := l.head
	for j:=1;j<i;j++ {
		next = next.next
	}
	// 现在next指针指向第i个节点
	return next, nil
}
