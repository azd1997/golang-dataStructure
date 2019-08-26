package linearlist

import (
	"log"
	"testing"
)

func TestSequenceList(t *testing.T) {

	// 测试初始化状态
	l := &SequenceList{}
	log.Println("初始化顺序表：", l)

	// 测试插入第一个元素
	if err := l.Insert(1, 3); err != nil {
		t.Error(err)
	}
	l.PrintAll()

	// 测试继续插入19个元素，使顺序表满
	for i:=0;i<19;i++ {
		// 插入1~19数字
		if err := l.Insert(i+2, i+1); err != nil {
			t.Error(err)
		}
	}
	l.PrintAll()
	// 继续插入一个1个元素，正常应该插不进
	if err := l.Insert(21, 21); err == nil {
		t.Error("超出MAXSIZE还能插入？！")
	}
	l.PrintAll()
	// 删除一个元素
	if err := l.Delete(7); err != nil {
		t.Error(err)
	}
	l.PrintAll()
	// 按索引查找
	elem1, err := l.FindByIndex(7)
	if  err != nil {
		t.Error(err)
	}
	log.Println("found elem: ", elem1)
	// 按值查找
	elems, err := l.FindByValue(3)
	if  err != nil {
		t.Error(err)
	}
	log.Println("found elems: ", elems)
}