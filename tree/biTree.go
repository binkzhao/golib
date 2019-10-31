package tree

import "container/list"

type biTree struct {
	root *biTreeNode // 根节点
}

// New一个二叉树
func NewBiTree(root *biTreeNode) *biTree {
	return &biTree{root: root}
}

//获得二叉树总结点数
func (this *biTree) GetSize() int {
	return this.root.size
}

//判断二叉树是否为空
func (this *biTree) IsEmpty() bool {
	return this.root != nil
}

//获得二叉树根节点
func (this *biTree) GetRoot() *biTreeNode {
	return this.root
}

//获得二叉树高度，根节点层为1
func (this *biTree) GetHeight() int {
	return this.root.height
}

//获得第一个与数据e相等的节点
func (this *biTree) Find(e interface{}) *biTreeNode {
	if this.root == nil {
		return nil
	}
	p := this.root
	return isEqual(e, p)
}

func isEqual(e interface{}, node *biTreeNode) *biTreeNode {
	if e == node.GetData() {
		return node
	}

	if node.HasLChild() {
		lp := isEqual(e, node.GetLChild())
		if lp != nil {
			return lp
		}
	}

	if node.HasRChild() {
		rp := isEqual(e, node.GetRChild())
		if rp != nil {
			return rp
		}

	}

	return nil
}

// 求第n层结点个数
func (this *biTree) GetKthNum(n int) int {
	return getKthNum(this.root, n)
}

func getKthNum(node *biTreeNode, n int) int {
	if node == nil || n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return getKthNum(node.lChild, n-1) + getKthNum(node.rChild, n-1)
}

// 求叶子结点个数
func (this *biTree) GetLeafNum() int {
	return getLeafNum(this.root)
}

func getLeafNum(node *biTreeNode) int {
	if node == nil {
		return 0
	}

	if !node.HasLChild() && !node.HasRChild() {
		return 1
	}

	return getLeafNum(node.lChild) + getLeafNum(node.rChild)
}

//先序递归遍历：根-> 左子树 -> 右子树，并保存在链表里
func (this *biTree) PreOrderTraverse() *list.List {
	l := list.New()
	preOrderTraverse(this.root, l)
	return l
}

func preOrderTraverse(rt *biTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	l.PushBack(rt)
	preOrderTraverse(rt.GetLChild(), l)
	preOrderTraverse(rt.GetRChild(), l)
}

//中序递归遍历：左子树-> 根 -> 右子树，并保存在链表里
func (this *biTree) InOrderTraverse() *list.List {
	l := list.New()
	inOrderTraverse(this.root, l)
	return l
}

func inOrderTraverse(rt *biTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	inOrderTraverse(rt.GetLChild(), l)
	l.PushBack(rt)
	inOrderTraverse(rt.GetRChild(), l)
}

//后序递归遍历：左子树-> 右子树 ->  根，并保存在链表里
func (this *biTree) PostOrderTraverse() *list.List {
	l := list.New()
	postOrderTraverse(this.root, l)
	return l
}

func postOrderTraverse(rt *biTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	postOrderTraverse(rt.GetLChild(), l)
	postOrderTraverse(rt.GetRChild(), l)
	l.PushBack(rt)
}
