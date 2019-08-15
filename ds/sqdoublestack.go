package ds

// 共享栈类型
type SqDoubleStack struct {
	Data [MaxSize]ElemType
	Top1 int // 栈1栈顶指针，为-1时表示空栈
	Top2 int // 栈1栈顶指针，为MaxSize时表示空栈
}

// 初始化栈
func NewSqDoubleStack() *SqDoubleStack {
	return &SqDoubleStack{
		Data: [MaxSize]ElemType{},
		Top1: -1,
		Top2: MaxSize,
	}
}

// 压栈
func (s *SqDoubleStack) Push(e ElemType, stackNumber int) bool {
	// 栈满
	if s.IsFull() {
		return false
	}

	// 获取新的栈顶指针
	var top int
	if stackNumber == 1 {
		s.Top1++
		top = s.Top1
	} else {
		s.Top2--
		top = s.Top2
	}

	s.Data[top] = e

	return true
}

// 出栈
func (s *SqDoubleStack) Pop(e *ElemType, stackNumber int) bool {
	// 栈空判断
	if s.IsEmpty(stackNumber) {
		return false
	}

	// 获取栈顶指针
	var top int
	if stackNumber == 1 {
		s.Top1--
		top = s.Top1
	} else {
		s.Top2++
		top = s.Top2
	}

	*e = s.Data[top]

	return true
}

// 栈空判断
func (s SqDoubleStack) IsEmpty(stackNumber int) bool {
	// 栈1
	if stackNumber == 1 && s.Top1 == -1 {
		return true
	}

	// 栈2
	if stackNumber == 2 && s.Top1 == MaxSize {
		return true
	}

	return false
}

// 栈满判断
func (s SqDoubleStack) IsFull() bool {
	if s.Top1+1 == s.Top2 {
		return true
	}

	return false
}

// 获取栈的元素个数
func (s SqDoubleStack) GetLength(stackNumber int) int {
	// 栈1
	if stackNumber == 1 {
		return s.Top1 + 1
	}

	// 栈2
	return MaxSize - s.Top2
}

// 获取栈顶元素
func (s SqDoubleStack) GetTop(e *ElemType, stackNumber int) bool {
	// 栈空
	if s.IsEmpty(stackNumber) {
		return false
	}

	var top int
	if stackNumber == 1 {
		top = s.Top1
	} else {
		top = s.Top2
	}

	*e = s.Data[top]

	return true
}
