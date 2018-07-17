//双链表
package linerList

import (
	"errors"
)

//双链表接口
type DoubleListInterface interface {
	IsEmpty() bool                        //判断链表是否为空链表
	Len() int                             //返回链表长度
	Insert(i int, elem interface{}) (*DuLNode, error) //指定位置插入元素
	Append(elem interface{}) *DuLNode              //链表末尾插入元素
	Get(i int) (*DuLNode, error)        //获取链表中第i个元素
	Remove(i int) (interface{}, error)    //删除链表中第i个元素,并返回Elem值
	Traverse(h traverseHandler) //遍历
}
type traversedHandler = func(k int, node *DuLNode)
//双链表节点
type DuLNode struct {
	Elem interface{}
	prev,next *DuLNode
}

type DoubleList struct {
	head *DuLNode //指向头节点的指针
}


//初始化一个单链表
func InitDoubleList() *DoubleList {
	head := &DuLNode{} //头结点
	return &DoubleList{head:head}
}



//判断链表是否为空
func (l *DoubleList ) IsEmpty() bool {
	return l.head.next == l.head.prev && l.head.next == l.head
}


//下一个节点
func (l *DuLNode) Next() *DuLNode {
	return l.next
}

//上一个节点
func (l *DuLNode) Prev() *DuLNode {
	return l.prev
}

//返回长度
func (l *DoubleList) Len() int {
	var length int
	var p *DuLNode
	p = l.head
	for p.Next() != nil {
		length++
		p = p.next
	}
	return length

}

//插入,从第i个位置插入一个元素 i范围 1 <= i <= 表长 + 1
//时间复杂度O(n)
func (l *DoubleList) Insert(i int, elem interface{})(*DuLNode, error) {
	var p *DuLNode
	var j int
	p = l.head
	j = 0
	for p != nil && j < i-1 { //寻找i-1个数据节点
		p = p.next
		j++
	}
	if p == nil || j > i-1 {
		return nil,errors.New("i小于1或者超出链表范围")
	}
	newNode := &DuLNode{Elem: elem,prev:p,next:p.next}
	if p.next != nil {//非尾部插入
		p.next.prev = newNode
	}

	p.next = newNode
	return newNode,nil
}

//尾部追加元素
//时间复杂度O(n)
func (l *DoubleList) Append(elem interface{}) *DuLNode {
	var p *DuLNode
	p = l.head
	for p.next != nil {
		p = p.next
	}
	newNode := &DuLNode{Elem: elem,prev:p}
	p.next = newNode
	return newNode
}

//返回第i个元素
func (l *DoubleList) Get(i int) (*DuLNode, error) {
	var p *DuLNode
	var j int
	p = l.head.next              //第一个数据节点
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
func (l *DoubleList) Remove(i int) (interface{}, error) {
	var p *DuLNode
	var j int
	p = l.head //队头指针
	j = 0 //队头下标
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

func (l *DoubleList) Traverse(h traversedHandler) {
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