package tree

import "math"

// 二叉树结点
type biTreeNode struct {
	data   interface{} // 数据域
	parent *biTreeNode // 父节点
	lChild *biTreeNode // 左孩子
	rChild *biTreeNode // 右孩子
	height int         // 以该结点为根的树的高度
	size   int         // 以该结点为根的树的节点数(包括结点本身)
}

// 创建一个结点
func NewBiTreeNode(data interface{}) *biTreeNode {
	return &biTreeNode{data: data, size: 1, height: 1}
}

//获得节点数据
func (this *biTreeNode) GetData() interface{} {
	if this == nil {
		return nil
	}
	return this.data
}

//设置节点数据
func (this *biTreeNode) SetData(data interface{}) {
	this.data = data
}

//判断是否有父亲
func (this *biTreeNode) HasParent() bool {
	return this.parent != nil
}

//获得父亲节点
func (this *biTreeNode) GetParent() *biTreeNode {
	if !this.HasParent() {
		return nil
	}
	return this.parent
}

//设置父亲节点
func (this *biTreeNode) setParent(p *biTreeNode) {
	this.parent = p
	this.parent.SetHeight() //更新父结点及其祖先高度
	this.parent.SetSize()   //更新父结点及其祖先规模
}

//断开与父亲的关系
func (this *biTreeNode) CutOffParent() {
	if !this.HasParent() {
		return
	}
	if this.IsLChild() {
		this.parent.lChild = nil //断开该节点与父节点的连接
	} else {
		this.parent.rChild = nil //断开该节点与父节点的连接
	}

	this.parent = nil       //断开该节点与父节点的连接
	this.parent.SetHeight() //更新父结点及其祖先高度
	this.parent.SetSize()   //更新父结点及其祖先规模
}

//判断是否有左孩子
func (this *biTreeNode) HasLChild() bool {
	return this.lChild != nil
}

//获得左孩子节点
func (this *biTreeNode) GetLChild() *biTreeNode {
	if !this.HasLChild() {
		return nil
	}
	return this.lChild
}

//设置当前结点的左孩子,返回原左孩子
func (this *biTreeNode) SetLChild(lc *biTreeNode) *biTreeNode {
	oldLC := this.lChild
	if this.HasLChild() {
		this.lChild.CutOffParent() //断开当前左孩子与结点的关系
	}
	if lc != nil {
		lc.CutOffParent() //断开lc与其父结点的关系
		this.lChild = lc  //确定父子关系
		lc.setParent(this)
		this.SetHeight() //更新当前结点及其祖先高度
		this.SetSize()   //更新当前结点及其祖先规模
	}
	return oldLC
}

//判断是否有右孩子
func (this *biTreeNode) HasRChild() bool {
	return this.rChild != nil
}

//获得右孩子节点
func (this *biTreeNode) GetRChild() *biTreeNode {
	if !this.HasRChild() {
		return nil
	}
	return this.rChild
}

//设置当前结点的右孩子,返回原右孩子
func (this *biTreeNode) SetRChild(rc *biTreeNode) *biTreeNode {
	oldRC := this.rChild
	if this.HasRChild() {
		this.rChild.CutOffParent() //断开当前左孩子与结点的关系
	}
	if rc != nil {
		rc.CutOffParent() //断开rc与其父结点的关系
		this.rChild = rc  //确定父子关系
		rc.setParent(this)
		this.SetHeight() //更新当前结点及其祖先高度
		this.SetSize()   //更新当前结点及其祖先规模
	}
	return oldRC
}

//判断是否为叶子结点
func (this *biTreeNode) IsLeaf() bool {
	return !this.HasLChild() && !this.HasRChild()
}

//判断是否为某结点的左孩子
func (this *biTreeNode) IsLChild() bool {
	return this.HasParent() && this == this.parent.lChild
}

//判断是否为某结点的右孩子
func (this *biTreeNode) IsRChild() bool {
	return this.HasParent() && this == this.parent.rChild
}

//取结点的高度,即以该结点为根的树的高度
func (this *biTreeNode) GetHeight() int {
	return this.height
}

//更新当前结点及其祖先的高度
func (this *biTreeNode) SetHeight() {
	newH := 1 //新高度初始化为1,高度等于左右子树高度加1中的大者
	if this.HasLChild() {
		newH = int(math.Max(float64(newH), float64(1+this.GetLChild().GetHeight())))
	}
	if this.HasRChild() {
		newH = int(math.Max(float64(newH), float64(1+this.GetRChild().GetHeight())))
	}
	if newH == this.height {
		//高度没有发生变化则直接返回
		return
	}

	this.height = newH //否则更新高度
	if this.HasParent() {
		this.GetParent().SetHeight() //递归更新祖先的高度
	}
}

//取以该结点为根的树的结点数
func (this *biTreeNode) GetSize() int {
	return this.size
}

//更新当前结点及其祖先的子孙数
func (this *biTreeNode) SetSize() {
	this.size = 1 //初始化为1,结点本身
	if this.HasLChild() {
		this.size += this.GetLChild().GetSize() //加上左子树规模
	}
	if this.HasRChild() {
		this.size += this.GetRChild().GetSize() //加上右子树规模
	}

	if this.HasParent() {
		this.parent.SetSize() //递归更新祖先的规模
	}
}
