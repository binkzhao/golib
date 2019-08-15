package ds

import "fmt"

// 带头结点的单链表

// 节点类型
type LinkNode struct {
	Data ElemType
	Next *LinkNode
}

// 单链表
type LinkList struct {
	Length int
	Next   *LinkNode
}

func NewLinkList() *LinkList {
	return &LinkList{
		Length: 0,
		Next: &LinkNode{
			Data: -1,
			Next: nil,
		},
	}
}

// 获取指定位置的元素
func (l LinkList) GetElem(i int, e *ElemType) bool {
	j := 0
	p := l.Next // p此时指向头结点
	for p.Next != nil && j < i {
		p = p.Next
		j++
	}

	if p == nil || j != i {
		return false
	}

	*e = p.Data

	return true
}

// 插入元素e到位置i
func (l *LinkList) Insert(i int, e ElemType) bool {
	j := 1
	p := l.Next
	for p.Next != nil && j < i {
		p = p.Next
		j++
	}

	if p == nil || j > i {
		return false
	}

	s := &LinkNode{Data: e, Next: nil}
	s.Next = p.Next
	p.Next = s

	l.Length++

	return true
}

// 获取长度
func (l LinkList) GetLength() int {
	return l.Length
}

// 删除位置i元素
func (l *LinkList) Delete(i int, e *ElemType) bool {
	j := 1
	p := l.Next
	//找到第i-1那个节点
	for p.Next != nil && j < i {
		p = p.Next
		j++
	}

	// 第i个元素不存在
	if p.Next == nil || j > i {
		return false
	}

	q := p.Next
	p.Next = q.Next
	*e = q.Data

	l.Length--

	return true
}

func (l *LinkList) Print() {
	p := l.Next // p此时为头结点
	if p.Next == nil {
		fmt.Printf("list.length =  %d, list is empty\n", l.Length)
		return
	}

	a := make([]ElemType, l.Length)
	i := 0
	for p.Next != nil && i < l.Length {
		a[i] = p.Next.Data
		p = p.Next
		i++
	}

	fmt.Println("list.length =", l.Length, ",list.data: ", a)
	return
}

// 头插入法把1到n个数据插入到单链表中
func CreateListHead(l *LinkList, n int) {
	for i := 1; i <= n; i++ {
		l.Insert(1, ElemType(i))
	}
}

// 尾插法把1到n个数据插入到单链表中
func CreateListTail(l *LinkList, n int) {
	for i := 1; i <= n; i++ {
		l.Insert(l.Length+1, ElemType(i))
	}
}
