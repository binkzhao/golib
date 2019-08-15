package ds

// 线性表顺序存储结构

type SqList struct {
	Data   [MaxSize]ElemType
	Length int
}

func NewSqList() *SqList {
	return &SqList{
		Data:   [MaxSize]ElemType{},
		Length: 0,
	}
}

// 获取指定位置的元素
func (l SqList) GetElem(i int, e *ElemType) bool {
	if l.Length == 0 || i < 1 || i > l.Length {
		return false
	}

	*e = l.Data[i-1]
	return true
}

// 插入元素
func (l *SqList) Insert(i int, e ElemType) bool {
	if l.Length == MaxSize {
		return false
	}

	if i < 1 || i > l.Length+1 {
		return false
	}

	// 若插入数据位置不再表尾，需要把插入位置后的数据元素全部向后移动一位
	if i < l.Length {
		for k := l.Length - 1; k >= i-1; k-- {
			l.Data[k+1] = l.Data[k]
		}
	}

	l.Data[i-1] = e
	l.Length++
	return true
}

// 获取长度
func (l SqList) GetLength() int {
	return l.Length
}

// 删除元素
func (l *SqList) Delete(i int, e *ElemType) bool {
	if l.Length == 0 || i < 1 || i > l.Length {
		return false
	}

	*e = l.Data[i-1]

	// 将删除位置后面的元素前移
	if i < l.Length {
		for k := i; k < l.Length; k++ {
			l.Data[k-1] = l.Data[k]
		}
	}

	l.Length--
	return true
}
