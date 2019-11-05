package tree

import (
	"fmt"
	"testing"
)

func TestBiTree(t *testing.T) {
	a := NewBiTreeNode(1)
	tree1 := NewBiTree(a)
	a.SetLChild(NewBiTreeNode(2))
	a.SetRChild(NewBiTreeNode(5))
	a.GetLChild().SetRChild(NewBiTreeNode(3))
	a.GetLChild().GetRChild().SetLChild(NewBiTreeNode(4))
	a.GetRChild().SetLChild(NewBiTreeNode(6))
	a.GetRChild().SetRChild(NewBiTreeNode(7))
	a.GetRChild().GetLChild().SetRChild(NewBiTreeNode(9))
	a.GetRChild().GetRChild().SetRChild(NewBiTreeNode(8))

	node2 := a.GetLChild()
	node9 := a.GetRChild().GetLChild().GetRChild()

	fmt.Println("结点2是叶子结点吗? ", node2.IsLeaf())
	fmt.Println("结点9是叶子结点吗? ", node9.IsLeaf())
	fmt.Println("这棵树的结点总数：", tree1.GetSize())

	l := tree1.InOrderTraverse() //中序遍历
	fmt.Println("中序遍历二叉树：")
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*biTreeNode)
		fmt.Printf("%v\t", obj.data)
	}
	result := tree1.Find(6)
	fmt.Printf("\n结点6的父节点数据：%v\n", result)
	fmt.Println("树的高度：", tree1.GetHeight(), tree1.root.rChild.rChild.GetHeight())
	fmt.Println("第三层结点的个数为：", tree1.GetKthNum(4))
	fmt.Println("树的总叶子结点的个数为：", tree1.GetLeafNum())

	// 层次遍历
	l = tree1.BreadthFirstTraverse()
	fmt.Println("层次遍历二叉树：")
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*biTreeNode)
		fmt.Printf("%v\t", obj.data)
	}
	fmt.Println("")
}
