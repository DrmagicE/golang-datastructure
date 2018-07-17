package linerList

import (
	"errors"
)

//单链表接口
type SingleListInterface interface {
	IsEmpty() bool                                  //判断链表是否为空链表
	Len() int                                       //返回链表长度
	Insert(i int, elem interface{}) (*LNode, error) //指定位置插入元素
	Append(elem interface{}) *LNode                 //链表末尾插入元素
	Get(i int) (*LNode, error)                      //获取链表中第i个元素
	Remove(i int) (interface{}, error)              //删除链表中第i个元素,并返回Elem值
	Traverse(h traverseHandler) //遍历
}

type traverseHandler = func(k int, node *LNode)

type SingleList struct {
	head *LNode //指向头节点的指针
}

//单链表数据节点与头节点结构体
type LNode struct {
	Elem interface{}
	next *LNode
}



//初始化一个单链表
func InitSingleList() *SingleList {
	head := &LNode{} //头结点
	return &SingleList{head: head}
}

//判断链表是否为空
func (l *SingleList) IsEmpty() bool {
	return l.head.next == nil
}

func (l *LNode) Next() *LNode {
	return l.next
}

func (l *SingleList) Len() int {
	var length int
	var p *LNode
	p = l.head
	for p.Next() != nil {
		length++
		p = p.next
	}
	return length

}

//插入,从第i个位置插入一个元素 i最小=1
//时间复杂度O(n)
func (l *SingleList) Insert(i int, elem interface{}) (*LNode, error) {
	var p *LNode
	var j int
	p = l.head
	j = 0
	for p != nil && j < i-1 { //寻找i-1个数据节点
		p = p.next
		j++
	}
	if p == nil || j > i-1 {
		return nil, errors.New("i小于1或者超出链表范围")
	}
	newNode := &LNode{Elem: elem}
	newNode.next = p.next
	p.next = newNode
	return newNode, nil
}

//尾部追加元素
//时间复杂度O(n)
func (l *SingleList) Append(elem interface{}) *LNode {
	var p *LNode
	p = l.head
	for p.next != nil {
		p = p.next
	}
	newNode := &LNode{Elem: elem}
	p.next = newNode
	return newNode
}

//返回第i个元素
func (l *SingleList) Get(i int) (*LNode, error) {
	var p *LNode
	var j int
	p = l.head.next         //第一个数据节点
	j = 1                   //数据节点下标
	for p != nil && j < i { //循环查找第i个节点
		p = p.next
		j++
	}
	if p == nil || j > i {
		return nil, errors.New("下标越界，元素不存在")
	}
	return p, nil
}

//删除并返回第i个元素
//时间复杂度O(n)
func (l *SingleList) Remove(i int) (interface{}, error) {
	var p *LNode
	var j int
	p = l.head //队头指针
	j = 0      //队头下标

	for p.next != nil && j < i-1 { //寻找第i个元素的前面一个元素
		p = p.next
		j++
	}
	if p.next == nil || j > i-1 { //下标越界
		return nil, errors.New("下标越界，元素不存在")
	}
	elem := p.next.Elem
	//free(p.next) 的操作GC来帮忙做了
	p.next = p.next.next
	return elem, nil
}

func (l *SingleList) Traverse(h traverseHandler) {
	var i int
	i = 1
	node, _ := l.Get(i)
	for node != nil {
		if h != nil {
			h(i, node)
		}
		i++
		node = node.Next()
	}
}
