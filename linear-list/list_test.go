package linearlist

import (
	"errors"
	"log"
	"testing"
)

func listTest(l LinearList) (err error) {
	// 测试空表
	log.Println("初始化顺序表：", l)

	// 空表状态测试按值查找
	elems := l.FindByValue(3)
	log.Println("found elems: ", elems)

	// 测试插入第一个元素
	if err := l.Insert(1, 3); err != nil {
		return err
	}
	l.PrintAll()

	// 测试继续插入19个元素，使顺序表满
	for i := 0; i < 9; i++ {
		// 插入1~9数字
		if err := l.Insert(i+2, i+1); err != nil {
			return err
		}
	}
	l.PrintAll()
	// 继续插入一个1个元素，正常应该插不进
	if err := l.Insert(11, 11); err == nil {
		return errors.New("没有空间了怎么还能插入？！！")
	}
	l.PrintAll()
	// 删除一个元素
	if err := l.Delete(7); err != nil {
		return err
	}
	l.PrintAll()
	// 按索引查找
	elem1, err := l.FindByIndex(7)
	if err != nil {
		return err
	}
	log.Println("found elem: ", elem1)
	// 按值查找
	elems = l.FindByValue(3)
	log.Println("found elems: ", elems)

	return nil
}

func TestSequenceList(t *testing.T) {

	l := &SequenceList{}
	// 这里要注意！！！实现LinearList接口的是*SequenceList!!!
	if err := listTest(l); err != nil {
		t.Error(err)
	}

	// // 测试初始化状态
	// l := &SequenceList{}
	// log.Println("初始化顺序表：", l)

	// // 空表状态测试按值查找
	// elems := l.FindByValue(3)
	// log.Println("found elems: ", elems)

	// // 测试插入第一个元素
	// if err := l.Insert(1, 3); err != nil {
	// 	t.Error(err)
	// }
	// l.PrintAll()

	// // 测试继续插入19个元素，使顺序表满
	// for i:=0;i<19;i++ {
	// 	// 插入1~19数字
	// 	if err := l.Insert(i+2, i+1); err != nil {
	// 		t.Error(err)
	// 	}
	// }
	// l.PrintAll()
	// // 继续插入一个1个元素，正常应该插不进
	// if err := l.Insert(21, 21); err == nil {
	// 	t.Error("超出MAXSIZE还能插入？！")
	// }
	// l.PrintAll()
	// // 删除一个元素
	// if err := l.Delete(7); err != nil {
	// 	t.Error(err)
	// }
	// l.PrintAll()
	// // 按索引查找
	// elem1, err := l.FindByIndex(7)
	// if  err != nil {
	// 	t.Error(err)
	// }
	// log.Println("found elem: ", elem1)
	// // 按值查找
	// elems = l.FindByValue(3)
	// log.Println("found elems: ", elems)
}

func TestSingleLinkedList(t *testing.T) {

	l := &SingleLinkedList{}
	// 这里要注意！！！实现LinearList接口的是*SequenceList!!!
	if err := listTest(l); err != nil {
		t.Error(err)
	}
}

func TestDoubleLinkedList(t *testing.T) {

	l := &DoubleLinkedList{}
	// 这里要注意！！！实现LinearList接口的是*SequenceList!!!
	if err := listTest(l); err != nil {
		t.Error(err)
	}
}

func TestCircularSingleLinkedList(t *testing.T) {

	l := &CircularSingleLinkedList{}
	// 这里要注意！！！实现LinearList接口的是*SequenceList!!!
	if err := listTest(l); err != nil {
		t.Error(err)
	}
}

func TestCircularDoubleLinkedList(t *testing.T) {

	l := &CircularDoubleLinkedList{}
	// 这里要注意！！！实现LinearList接口的是*SequenceList!!!
	if err := listTest(l); err != nil {
		t.Error(err)
	}
}

