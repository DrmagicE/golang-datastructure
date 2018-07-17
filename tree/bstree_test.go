package tree

import (
	"testing"
	"dataStructure/queue"
	"strconv"
)

//测试二叉查找树
func TestBitTree_SearchBST(t *testing.T) {
	var p []interface{}
	p = []interface{}{
		NewBSTElem(45,45),
		NewBSTElem(12,12),
		NewBSTElem(3,3),nil,nil,
		NewBSTElem(37,37),
		NewBSTElem(24,24),
		nil,nil,nil,
		NewBSTElem(53,53),
		nil,
		NewBSTElem(100,100),
		NewBSTElem(61,61),
		nil,
		NewBSTElem(90,90),
		NewBSTElem(78,78),
		nil,nil,nil,nil,
	}
	q :=queue.CreateQueue(p...)
	tree := CreateBitTree(q)
	tree.InsertBST(NewBSTElem(2,2))
	tree.InsertBST(NewBSTElem(4,4))
	tree.InsertBST(NewBSTElem(55,55))
	var str string
	tree.PostOrderTraverse(func(node BitNode) {
		str += strconv.Itoa(node.Elem.(IBSTElem).GetElem().(int))
	})
	if str != "243243712557890611005345" {
		t.Error("二叉查找树插入失败，期望:",str,"实际",str)
	}
}

//测试删除二叉排序树
func TestBitTree_DeleteBST(t *testing.T) {
	var p []interface{}
	p = []interface{}{
		NewBSTElem(45,45),
		NewBSTElem(12,12),
		NewBSTElem(3,3),nil,nil,
		NewBSTElem(37,37),
		NewBSTElem(24,24),
		nil,nil,nil,
		NewBSTElem(53,53),
		nil,
		NewBSTElem(100,100),
		NewBSTElem(61,61),
		nil,
		NewBSTElem(90,90),
		NewBSTElem(78,78),
		nil,nil,nil,nil,
	}
	q :=queue.CreateQueue(p...)
	tree := CreateBitTree(q)
	tree.InsertBST(NewBSTElem(2,2))
	tree.InsertBST(NewBSTElem(4,4))
	tree.InsertBST(NewBSTElem(55,55))

	if !tree.DeleteBST(61) {
		t.Error("应该删除成功，但却失败")
	}

	var str string
	tree.InOrderTraverse(func(node BitNode) {
		str += strconv.Itoa(node.Elem.(IBSTElem).GetElem().(int))
	})
	if str != "2341224374553557890100" {
		t.Error("删除错误，期望:2341224374553557890100","实际:",str)
	}

}
