//二叉排序树
package tree

import (
	"errors"
)

//排序树判断大小
type IBSTElem interface {
	Key() int             //比较的值
	GetElem() interface{} //获取元素值
}

type BSTElem struct {
	key  int
	Elem interface{}
}

func (e *BSTElem) Key() int {
	return e.key
}

func (e *BSTElem) GetElem() interface{} {
	return e.Elem
}

func NewBSTElem(key int, elem interface{}) *BSTElem {
	return &BSTElem{key: key, Elem: elem}
}

//查找成功，返回对应节点
//查找不成功，返回最后查抄的节点，并产生error
func (t *BitTree) SearchBST(key int, parrent *BitNode) (*BitNode, error) {
	var node *BitNode
	node = t
	if t.IsEmpty() { //查空了都没查到
		return parrent, errors.New("not found")
	}
	e, ok := t.Elem.(IBSTElem)
	if !ok {
		panic("二叉排序树节点类型不符")
	}
	if e.Key() == key { //相等，直接返回
		return node, nil
	} else if e.Key() < key { //从右子树开始
		return t.rchild.SearchBST(key, t)
	} else { //从左子树开始
		return t.lchild.SearchBST(key, t)
	}
}

//插入元素 是否插入成功
func (t *BitTree) InsertBST(elem IBSTElem) bool {
	if t.IsEmpty() {
		return false
	}
	if parentNode, err := t.SearchBST(elem.Key(), nil); err != nil { //查找不成功
		node := &BitNode{Elem: elem}
		if parentNode == nil { //空树，直接插入根节点
			t = node
		} else if parentNode.Elem.(IBSTElem).Key() < elem.Key() { //右边插入
			parentNode.rchild = node
		} else { //左边插入
			parentNode.lchild = node
		}
		return true
	}
	return false
}

//删除元素
func (t *BitTree) DeleteBST(key int) bool {
	node, err := t.SearchBST(key, nil)
	if err == nil { //找到对应元素
		if node.rchild == nil { //右子树空
			node = node.lchild
		} else if node.lchild == nil { //左子树为空
			node = node.rchild
		} else { //左右子树都不为空,1.采用中序遍历的直接前驱替换，2.删掉直接前驱
			var pre *BitNode //直接前驱
			pre = node.lchild
			var q *BitNode //直接前驱的双亲节点
			q = node
			for pre.rchild != nil { //寻找直接前驱
				q = pre //直接前驱的双亲节点
				pre = pre.rchild
			}

			node.Elem = pre.Elem //1.直接前驱替换掉被删除节点

			//2.开始删除直接前驱节点
			if q != node { //q是直接前驱的双亲节点
				q.rchild = pre.lchild //重新接上直接前驱的左子树作为q的右子树（此时直接前驱右子树必然是nil, 因为前面的 =》 for pre != nil {}）
			} else {
				q.lchild = pre.lchild //直接前驱就是被删除节点的左节点的情况
			}
			return true
		}
	}
	return false
}
