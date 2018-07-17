//go不支持指针运算，只好用链栈代替
package stack

import (
	"errors"
)

type StackInterface interface {
	IsEmpty() bool             //判断是否空栈
	Len() int                  //返回栈长度
	Push(interface{})          //入栈
	Pop() (interface{}, error) //出栈
	GetTop() interface{}       //获取栈顶元素
}

type traverseHandler = func(k int,elem *SElem)

type Stack struct {
	top  *SElem
	size int
}

type SElem struct {
	Elem interface{}
	next *SElem
}

//初始化栈
func InitStack() *Stack {
	//初始化栈顶
	top := &SElem{Elem: nil, next: nil}
	return &Stack{top: top, size: 0}
}

//判空
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

//入栈
func (s *Stack) Push(elem interface{}) {
	s.top.Elem = elem
	e := &SElem{Elem: nil, next: s.top} //新栈顶
	s.top = e
	s.size++
}

//出栈
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("栈空")
	}
	p := s.top.next //栈顶元素
	s.top.next = p.next
	s.size--
	return p.Elem, nil

}

//获取栈顶元素
func (s *Stack) GetTop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.top.next.Elem
}

//栈长度
func (s *Stack) Len() int {
	return s.size
}

//遍历
func (s *Stack) Traverse(h traverseHandler) {
	var i int
	i = 1
	node := s.top.next
	for node != nil {
		if h != nil {
			h(i, node)
		}
		node = node.next
		i++
	}
}