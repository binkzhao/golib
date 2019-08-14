package datastruct

// 顺序存储的栈
type SqStack struct {
	Data [MaxSize]ElemType
	Top  int
}

// 初始化栈
func NewSqStack() *SqStack {
	return &SqStack{
		Data: [MaxSize]ElemType{},
		Top:  -1,
	}
}

// 压栈
func (s *SqStack) Push(e ElemType) bool {
	// 栈满
	if s.IsFull() {
		return false
	}

	s.Top++
	s.Data[s.Top] = e

	return true
}

// 出栈
func (s *SqStack) Pop(e *ElemType) bool {
	// 栈空判断
	if s.IsEmpty() {
		return false
	}

	*e = s.Data[s.Top]
	s.Top--

	return true
}

// 栈空判断
func (s SqStack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}

	return false
}

// 栈满判断
func (s SqStack) IsFull() bool {
	if s.Top == MaxSize-1 {
		return true
	}

	return false
}

// 获取栈的元素个数
func (s SqStack) GetLength() int {
	return s.Top + 1
}

func (s SqStack) GetTop(e *ElemType) bool {
	if s.IsEmpty() {
		return false
	}

	*e = s.Data[s.Top]

	return true
}
