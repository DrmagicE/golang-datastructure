package queue

import (
	"errors"
)

type QueueInterface interface {
	IsEmpty() bool             //判断是否空栈
	Len() int                  //返回栈长度
	GetHead() *QNode	//返回队头
	EnQueue(elem interface{}) //入队
	DeQueue() (interface{},error)//出队
	Traverse(fn traverseHandler) //遍历

}

type QNode struct {
	Elem interface{}
	next *QNode
}

type Queue struct {
	front, rear *QNode //队头，队尾指针
	size        int    //队列长度
}

type traverseHandler = func(k int, node *QNode)

//初始化一个空队列
func InitQueue() *Queue {
	front := &QNode{}
	rear := front
	return &Queue{front: front, rear: rear}
}

func CreateQueue(data ... interface{}) *Queue{
	q := InitQueue()
	for _,v:=range data {
		q.EnQueue(v)
	}

	return q
}

//判空
func (q *Queue) IsEmpty() bool {
	if q == nil {
		return true
	}
	return q.rear == q.front
}

func (q *Queue) Len() int {
	return q.size
}

//返回头元素
func (q *Queue) GetHead() *QNode {
	return q.front.next
}

//入队
func (q *Queue) EnQueue(data interface{}) {
	node := &QNode{Elem: data}
	q.rear.next = node
	q.rear = node
	q.size++
}

//出队
func (q *Queue) DeQueue() (interface{},error) {
	if q.IsEmpty() {
		return nil,errors.New("空队列")
	}
	p := q.front.next
	data := p.Elem
	q.front.next = p.next
	if q.rear == p { //删除的是队尾元素
		q.rear = q.front
	}
	q.size--
	return data,nil
}

//遍历队列
func (q *Queue) Traverse(fn traverseHandler) {
	node := q.front.next
	var i int
	i = 1
	for node != nil {
		fn(i, node)
		node = node.next
		i++
	}
}

