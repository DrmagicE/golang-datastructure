package tree

import (
	"testing"
	"math/rand"
	"time"
)

//测试插入
func TestAVLTree_InsertAVL(t *testing.T) {

	var tr *AVLTree
 	rand.Seed(time.Now().Unix())

	tr =InitAVLTree(NewBSTElem(0,0))
	for i:=0;i<10000 ;i++{
		tr,_ = tr.InsertAVL(NewBSTElem(rand.Int(),rand.Int()))
	}
	var prvSet bool
	var prvKey int
	tr.InOrderTraverse(func(node AVLNode) {
		if prvKey > node.Elem.Key() && prvSet == true{
			t.Error("avl树插入错误，无序")
		}
		if node.bf != EH && node.bf != LH && node.bf != RH {
			t.Error("avl树插入错误，平衡因子不合法")
		}
		prvKey = node.Elem.Key()
		prvSet = true
	})

}
