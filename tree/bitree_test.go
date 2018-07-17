package tree

import (
	"dataStructure/queue"
	"testing"
)

//先序遍历测试
func TestBitTree_PreOrderTraverse(t *testing.T) {
	var tests = []struct {
		input []interface{}
		want  string
	}{
		{input: []interface{}{"a", "b", "d", nil, nil, "e", nil, nil, "c", "f", nil, nil, "g"}, want: "abdecfg"},
		{input: []interface{}{"a", "b", "d", nil, nil, nil, nil}, want: "abd"},
		{input: []interface{}{"-", "+", "a", nil, nil, "*", "b", nil, nil, "-", "c", nil, nil, "d", nil, nil, "/", "e", nil, nil, "f", nil, nil}, want: "-+a*b-cd/ef"},
	}
	for k, v := range tests {
		q := queue.CreateQueue(v.input...)
		tree := CreateBitTree(q)
		var treeStr string
		treeStr = ""
		tree.PreOrderTraverse(func(node BitNode) {
			treeStr += node.Elem.(string)
		})
		if treeStr != v.want {
			t.Error("先序遍历测试失败", k)
		}
	}
}

//中序遍历测试
func TestBitTree_InOrderTraverse(t *testing.T) {
	var tests = []struct {
		input []interface{}
		want  string
	}{
		{input: []interface{}{"a", "b", "d", nil, nil, "e", nil, nil, "c", "f", nil, nil, "g"}, want: "dbeafcg"},
		{input: []interface{}{"a", "b", "d", nil, nil, nil, nil}, want: "dba"},
		{input: []interface{}{"-", "+", "a", nil, nil, "*", "b", nil, nil, "-", "c", nil, nil, "d", nil, nil, "/", "e", nil, nil, "f", nil, nil}, want: "a+b*c-d-e/f"},
	}
	for k, v := range tests {
		q := queue.CreateQueue(v.input...)
		tree := CreateBitTree(q)
		var treeStr string
		treeStr = ""
		tree.InOrderTraverse(func(node BitNode) {
			treeStr += node.Elem.(string)
		})
		if treeStr != v.want {
			t.Error("中序遍历测试失败", k)
		}
	}
}

//后续遍历测试
func TestBitTree_PostOrderTraverse(t *testing.T) {
	var tests = []struct {
		input []interface{}
		want  string
	}{
		{input: []interface{}{"a", "b", "d", nil, nil, "e", nil, nil, "c", "f", nil, nil, "g"}, want: "debfgca"},
		{input: []interface{}{"a", "b", "d", nil, nil, nil, nil}, want: "dba"},
		{input: []interface{}{"-", "+", "a", nil, nil, "*", "b", nil, nil, "-", "c", nil, nil, "d", nil, nil, "/", "e", nil, nil, "f", nil, nil}, want: "abcd-*+ef/-"},
	}
	for k, v := range tests {
		q := queue.CreateQueue(v.input...)
		tree := CreateBitTree(q)
		var treeStr string
		treeStr = ""
		tree.PostOrderTraverse(func(node BitNode) {
			treeStr += node.Elem.(string)
		})
		if treeStr != v.want {
			t.Error("后序遍历测试失败", k)
		}
	}
}

//测试层序遍历
func TestBitTree_LevelOrderTraverse(t *testing.T) {
	var tests = []struct {
		input []interface{}
		want  string
	}{
		{input: []interface{}{"a", "b", "d", nil, nil, "e", nil, nil, "c", "f", nil, nil, "g"}, want: "abcdefg"},
		{input: []interface{}{"a", "b", "d", nil, nil, nil, nil}, want: "abd"},
		{input: []interface{}{"-", "+", "a", nil, nil, "*", "b", nil, nil, "-", "c", nil, nil, "d", nil, nil, "/", "e", nil, nil, "f", nil, nil}, want: "-+/a*efb-cd"},
	}
	for _, v := range tests {
		q := queue.CreateQueue(v.input...)
		tree := CreateBitTree(q)
		var treeStr string
		treeStr = ""
		tree.LevelOrderTraverse(func(node BitNode) {
			treeStr += node.Elem.(string)
		})
		if treeStr != v.want {
			t.Error("层序遍历失败，期望：",v.want,"实际：",treeStr )
		}
	}
}

func isInsertFail(tree *BitTree, err error, success bool, want string, t *testing.T) {
	if err == nil { //插入成功
		if success {
			var str string
			tree.PreOrderTraverse(func(node BitNode) {
				str += node.Elem.(string)
			})
			if str != want {
				t.Error("不通过,需要", want, "实际", str)
			}
		} else {
			t.Error("应当插入失败，但却成功")
		}
	} else {
		if success {
			t.Error("应当插入成功，却失败", err, want)
		}
	}
}

//测试插入
func Test_InsertChild(t *testing.T) {
	var tests = []struct {
		initT     []interface{} //被插入的插入
		insertedT []interface{} //插入的树
		success   bool          //是否能插入成功
		leftWant  string        //如果插入左子树成功，插入后先序遍历的字符串
		rightWant string        //如果插入右子树成功，插入后先序遍历的字符串

	}{
		{ //插入成功
			initT:     []interface{}{"-", "+", "a", nil, nil, "*", "b", nil, nil, "-", "c", nil, nil, "d", nil, nil, "/", "e", nil, nil, "f", nil, nil},
			insertedT: []interface{}{"1", "2", "3", nil, "5", nil, nil},
			success:   true,
			leftWant:  "-+a*1235b-cd/ef",
			rightWant: "-+a*b1235-cd/ef",
		},
		{ //插入失败
			initT:     []interface{}{"a", "b", "d", nil, nil, "e", nil, nil, "c", "f", nil, nil, "g"},
			insertedT: []interface{}{"1", "2", "3", nil, "5", nil, nil, nil, "7", nil, nil},
			success:   false,
		},
	}
	var err error
	for _, v := range tests {
		var initTree *BitTree //需要被插入节点的树
		var initQ *queue.Queue
		var insertedTree *BitTree //被插入的树
		var insertedQ *queue.Queue
		var root *BitNode //需要被插入树的root节点
		var pos *BitNode  //插入的节点位置

		//测试左节点插入
		initQ = queue.CreateQueue(v.initT...)
		initTree = CreateBitTree(initQ)
		insertedQ = queue.CreateQueue(v.insertedT...)
		insertedTree = CreateBitTree(insertedQ)
		root, _ = initTree.Root()
		pos = root.Left().Right()
		err = initTree.InsertChild(pos, insertedTree, LEFT)
		isInsertFail(initTree, err, v.success, v.leftWant, t)

		//测试右节点插入
		initQ = queue.CreateQueue(v.initT...)
		initTree = CreateBitTree(initQ)
		insertedQ = queue.CreateQueue(v.insertedT...)
		insertedTree = CreateBitTree(insertedQ)
		root, _ = initTree.Root()
		pos = root.Left().Right()
		err = initTree.InsertChild(pos, insertedTree, RIGHT)
		isInsertFail(initTree, err, v.success, v.rightWant, t)

	}
}

