package ds

// 元素类型
type ElemType int

// 存储空间初始化分配量
const MaxSize = 20

// 线性表list接口类型
type Lister interface {
	GetElem(i int, e *ElemType) bool // 获取指定位置的数据元素
	Insert(i int, e ElemType) bool   // 指定位置插入元素
	Delete(i int, e *ElemType) bool  // 删除指定位置元素
	GetLength() int                  // 获取list的长度
}

// 栈stack接口类型
type Stacker interface {
	IsEmpty() bool           // 判断栈空
	Push(e ElemType) bool    // 压栈
	Pop(e *ElemType) bool    // 出栈
	GetTop(e *ElemType) bool // 获取栈订元素
	GetLength() int          // 获取栈的元素个数
	IsFull() bool            // 栈满判断
}

// 共享栈接口类型
type DoubleStacker interface {
	IsEmpty(stackNumber int) bool             // 判断栈空
	Push(e ElemType, stackNumber int) bool    // 压栈
	Pop(e *ElemType, stackNumber int) bool    // 出栈
	GetTop(e *ElemType, stackNumber int) bool // 获取栈订元素
	GetLength(stackNumber int) int            // 获取栈的元素个数
	IsFull() bool                             // 栈满判断
}
