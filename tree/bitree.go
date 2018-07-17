package tree

import (
	"errors"
	"dataStructure/queue"
)

const (
	LEFT = iota
	RIGHT
)

type BitTreeInterface interface {
	Root() *BitNode                                      //返回根节点
	/*Parent(node *BitNode) *BitNode                       //返回双亲节点
	LeftChild(node *BitNode) *BitNode                    //返回左孩子节点
	RightChild(node *BitNode) *BitNode                   //返回右孩子节点
	LeftSibling(node *BitNode) *BitNode                  //返回左兄弟
	RightSibling(node *BitNode) *BitNode                 //返回右兄弟*/
	InsertChild(node *BitNode, elem interface{}, lr int) //插入节点(左或右)
	//RemoveChild(node *BitNode,lr int)  interface{} //删除节点(左或右)
	PreOrderTraverse()                                   //先序遍历
	InOrderTraverse()                                    //中序遍历
	PostOrderTraverse()                                  //后序遍历
	LevelOrderTraverse()                                 //层序遍历
}

type traverseFn = func(node BitNode)

type BitTree = BitNode

//二叉树节点
type BitNode struct {
	Elem           interface{}
	lchild, rchild *BitNode //左右孩子
}



//左孩子
func (t *BitNode) Left()  *BitNode {
	return t.lchild;
}

//右孩子
func (t *BitNode) Right()  *BitNode {
	return t.rchild;
}

//返回跟节点
func (t *BitTree) Root () (*BitNode,error)  {
	if t.IsEmpty() {
		return nil,errors.New("empty")
	}
	return t,nil
}


//初始化二叉树
func InitBitTree(rootElem interface{}) *BitTree {
	node := &BitNode{Elem: rootElem}
	return node
}

//判空
func (t *BitTree) IsEmpty() bool {
	return t == nil
}

//左节点插入
func (t *BitTree) insertLeftChild(p *BitNode, c *BitTree) error {
	l := p.lchild
	c.rchild = l
	p.lchild = (*BitNode)(c)
	return nil
}
//右节点插入
func (t *BitTree) insertRightChild(p *BitNode, c *BitTree) error {
	r := p.rchild
	c.rchild = r
	p.rchild = (*BitNode)(c)

	return nil
}
// 初始条件: 二叉树T存在,p是T中某个结点的值,LR为0或1,非空二叉树c与T 不相交且右子树为空
// 操作结果: 根据LR为0或1,插入c为T中p结点的左或右子树。p结点的原有左或 右子树则成为c的右子树
func (t *BitTree) InsertChild(p *BitNode, c *BitTree, lr int) error {
	if c.rchild != nil {
		return errors.New("被插入的树右子树不为空")
	}
	switch lr {
	case LEFT:
		return t.insertLeftChild(p, c)
	case RIGHT:
		return t.insertRightChild(p,c)
	default:
		return errors.New("不合法的位置")
	}
}


//先序创建二叉树
func CreateBitTree(queue *queue.Queue) *BitTree {
	t,_ := queue.DeQueue()
	if t != nil {
		tree := InitBitTree(t)                         //根节点
		tree.lchild = CreateBitTree(queue) //先生成左树
		tree.rchild = CreateBitTree(queue) //再生成右树
		return tree
	} else {
		return nil
	}
}

//先序递归遍历
//1.访问根节点
//2.先序遍历左子树
//3.先序遍历右子树
func (t *BitTree)PreOrderTraverse(fn traverseFn) {
	if !t.IsEmpty() {
		fn(*t)
		t.lchild.PreOrderTraverse(fn)
		t.rchild.PreOrderTraverse(fn)
	}
}
//中顺递归遍历
//1.中序遍历左子树
//2.访问根节点
//3.中序遍历右子树
func (t *BitTree)InOrderTraverse(fn traverseFn) {
	//fmt.Println(t)
	if !t.IsEmpty() {
		t.lchild.InOrderTraverse(fn)
		fn(*t)
		t.rchild.InOrderTraverse(fn)
	}
}

//后序遍历
//1.后序遍历左子树
//2.后序遍历右子树
//3.访问根节点
func (t *BitTree) PostOrderTraverse(fn traverseFn) {
	if !t.IsEmpty() {
		t.lchild.PostOrderTraverse(fn)
		t.rchild.PostOrderTraverse(fn)
		fn(*t)
	}
}

//层序遍历
//每层从左到右遍历
func (t *BitTree) LevelOrderTraverse (fn traverseFn) {
	if !t.IsEmpty() {
		var node *BitNode
		q := queue.InitQueue()
		node,_ = t.Root() //根节点
		for node != nil{
			fn(*node)
			lt := node.lchild //左节点
			rt := node.rchild //右节点
			if lt != nil {
				q.EnQueue(lt)
			}
			if rt != nil {
				q.EnQueue(rt)
			}
			qNode,_ := q.DeQueue()
			if qNode != nil {
				node = qNode.(*BitNode)
			} else { //队列为空,遍历结束
				break;
			}
		}
	}
}
