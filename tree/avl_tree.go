//平衡二叉树

package tree

type AvlTreeinterface interface {

}

type AVLNode struct {
	Elem           IBSTElem
	bf             int      //平衡因子
	lchild, rchild *AVLNode //左右孩子
}

const (
	LH = 1  //左高
	EH = 0  //等高
	RH = -1 //右高
)

type AVLTree = AVLNode

//插入节点过程中，不平衡可能出现在下面4中情况中：
//1.对节点的左儿子的左子树进行一次插入  LL型插入
//2.对节点的左儿子的右子树进行一次插入  LR型插入
//3.对节点的右儿子的左子树进行一次插入  RL型插入
//4.对节点的右儿子的右子树进行一次插入  RR型插入

//LL型插入
//以t为根节点的 右旋处理
func RRotate(t *AVLTree) *AVLTree {
	lc := t.lchild //左孩子作为新的根节点
	t.lchild = lc.rchild
	lc.rchild = t
	return lc


}

//RR型插入
//以t为根的 左旋处理
func LRotate(t *AVLTree) *AVLTree {
	rc := t.rchild
	t.rchild = rc.lchild
	rc.lchild = t
	return rc
}

//用一个根节点初始化avl平衡树
func InitAVLTree(root IBSTElem) *AVLTree {
	node := &AVLNode{Elem: root, bf: 0}
	return node
}

//插入
/*func InsertAVL(t *AVLTree, elem IBSTElem) (*AVLTree,bool,bool){
	return insertAVL(t, elem)

}*/

//判空
func (t *AVLTree) IsEmpty() bool {
	return t == nil
}


//内部递归实现
//t是被插入的节点
func insertAVL (t *AVLTree, elem IBSTElem) (node *AVLTree,ok bool,taller bool) {
	if t == nil { //插入节点
		return &AVLNode{Elem:elem},true,true
	}
	if elem.Key() == t.Elem.Key() {
		return t,false,false //相同不插入
	}
	if elem.Key() < t.Elem.Key() { //左子树插入
		t.lchild,ok,taller = insertAVL(t.lchild,elem)
		if !ok { //插入不成功
			return t,ok,taller
		}
		if taller { //插入成功且树增高
			switch t.bf {
			case LH: //插入之前左树高，左左插入 或者是 左右插入，左平衡处理
				t = LeftBalance(t)
				return t,true, false
			case EH: //本来是等高的
				t.bf = LH
				return t,true, true
			case RH: //右边高
				t.bf = EH
				return t,true, false
			}
		}
	} else { //右子树插入
		t.rchild,ok, taller = insertAVL(t.rchild, elem)
		if !ok { //未插入成功
			return t,ok, taller
		}
		if taller { //插入成功且树增高
			switch t.bf {
			case LH: //插入之前左树高，现在等高
				t.bf = EH
				return t,true, false
			case EH: //本来是等高的
				t.bf = RH
				return t,true, true
			case RH: //右边高
				t.bf = EH
				t = RightBalance(t)
				return t,true, false  //旋转之后就不再增高了
			}
		}
	}
	return t,true,false
}

//插入AVL节点
func (t *AVLTree ) InsertAVL (elem IBSTElem) (*AVLTree,bool){
	if t.IsEmpty() {
		return t,false //插入失败
	}
	nt,ok,_:=insertAVL(t,elem)
	return nt,ok
}



//左平衡
func LeftBalance(t *AVLTree) *AVLTree {
	lc := t.lchild //左子树根节点
	switch lc.bf {
	case LH: //左左插入，右旋即可
		t.bf = EH
		lc.bf = EH
		t = RRotate(t)

	case RH: //左右插入，双旋
		rd := lc.rchild
		switch rd.bf { //修改跟节点以及其左孩子平衡因子，这部分比较难理解.
		case LH: //右节点左高
			t.bf = RH
			lc.bf = EH
		case EH: //等高
			t.bf = EH
			lc.bf = EH
		case RH: //右节点右高
			t.bf = EH
			lc.bf = LH
		}
		rd.bf = EH
		t.lchild = LRotate(t.lchild)
		t = RRotate(t)
	}

	return t
}
//右平衡
func RightBalance(t *AVLTree) *AVLTree{
	rc := t.rchild; //右子树根节点
	switch rc.bf {
	case RH: //右右插入 ，要作单左旋处理
		t.bf = EH
		rc.bf = EH
		t = LRotate(t)
	case LH: //右左插入，双旋
		ld := rc.lchild
		switch ld.bf { //修改跟节点以及其左孩子平衡因子，这部分比较难理解.
		case RH: //右节点左高  插入的是根节点的右节点
			t.bf = LH
			rc.bf = EH
		case EH: //等高      //插入作为rc的左根节点
			t.bf = EH
			rc.bf = EH
		case LH: //左节点高  //插入的是根节点的左节点
			t.bf = EH
			rc.bf = RH
		}
		ld.bf = EH
		t.rchild = RRotate(t.rchild)
		t = LRotate(t)
	}
	return t
}

type avlTraverseFn = func(node AVLNode)

func (t *AVLTree)InOrderTraverse(fn avlTraverseFn) {
	//fmt.Println(t)
	if !t.IsEmpty() {
		t.lchild.InOrderTraverse(fn)
		fn(*t)
		t.rchild.InOrderTraverse(fn)
	}
}
