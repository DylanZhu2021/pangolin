package rbTree

import (
	"fmt"
	"strings"
)

const (
	RED   bool = true
	BLACK bool = false
	Desc  bool = false
)

type Node struct {
	key   interface{}
	value interface{}

	left  *Node
	right *Node
	//parent *Node

	color bool
}

type RedBlackTree struct {
	size int
	root *Node
}

func NewNode(key, val interface{}) *Node {
	// 默认添加红节点
	return &Node{
		key:   key,
		value: val,
		left:  nil,
		right: nil,
		//parent: nil,
		color: RED,
	}
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

// key 只支持 rune int string
func compare(a, b interface{}) int {

	switch a.(type) {
	case string:
		na := a.(string)
		nb := b.(string)
		if Desc {
			return -strings.Compare(strings.ToLower(na), strings.ToLower(nb))
		}
		return strings.Compare(strings.ToLower(na), strings.ToLower(nb))
	case int:
		na := a.(int)
		nb := b.(int)
		if Desc {
			return nb - na
		}
		return na - nb
	}
	return int(a.(rune) - b.(rune))
}
func (n *Node) IsRed() bool {
	if n == nil {
		return BLACK
	}
	return n.color
}

func (tree *RedBlackTree) GetTreeSize() int {
	return tree.size
}

//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (n *Node) leftRotate() *Node {
	// 左旋转
	retNode := n.right
	n.right = retNode.left

	retNode.left = n
	retNode.color = n.color
	n.color = RED

	return retNode
}

//     node                    x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func (n *Node) rightRotate() *Node {
	//右旋转
	retNode := n.left
	n.left = retNode.right

	retNode.right = n
	retNode.color = n.color
	n.color = RED

	return retNode
}

// 颜色变换
func (n *Node) flipColors() {
	n.color = RED
	n.left.color = BLACK
	n.right.color = BLACK
}

// 维护红黑树
func (n *Node) updateRedBlackTree(isAdd int) *Node {
	// isAdd=0 说明没有新节点，无需维护
	if isAdd == 0 {
		return n
	}

	// 需要维护
	if n.right.IsRed() == RED && n.left.IsRed() != RED {
		n = n.leftRotate()
	}

	// 判断是否为情形3，是需要右旋转
	if n.left.IsRed() == RED && n.left.left.IsRed() == RED {
		n = n.rightRotate()
	}

	// 判断是否为情形4，是需要颜色翻转
	if n.left.IsRed() == RED && n.right.IsRed() == RED {
		n.flipColors()
	}

	return n
}

// 递归写法:向树的root节点中插入key,val
// 返回1, 代表加了节点
// 返回0, 代表没有添加新节点, 只更新key对应的value值
func (n *Node) add(key, val interface{}) (int, *Node) {
	if n == nil { // 默认插入红色节点
		return 1, NewNode(key, val)
	}

	isAdd := 0
	if compare(key, n.key) < 0 {
		isAdd, n.left = n.left.add(key, val)
	} else if compare(key, n.key) > 0 {
		isAdd, n.right = n.right.add(key, val)
	} else {
		// 对value值更新,节点数量不增加,isAdd = 0
		n.value = val
	}

	// 维护红黑树
	n = n.updateRedBlackTree(isAdd)

	return isAdd, n
}
func (tree *RedBlackTree) Get(key interface{}) interface{} {
	p := tree.root
	for p != nil {
		cmp := compare(key, p.key)
		if cmp < 0 {
			p = p.left
		} else if cmp > 0 {
			p = p.right
		} else {
			return p.value
		}
	}
	return nil
}

func (tree *RedBlackTree) Add(key, val interface{}) {
	if key == nil || val == nil {
		panic("加入nil")
	}
	isAdd, nd := tree.root.add(key, val)
	tree.size += isAdd
	tree.root = nd
	tree.root.color = BLACK //根节点为黑色节点
}

// 前序遍历打印出key,val,color
func (tree *RedBlackTree) PrintPreOrder() {
	resp := make([][]interface{}, 0)
	tree.root.printPreOrder(&resp)
	fmt.Println(resp)
}
func (tree *RedBlackTree) GetMin() (interface{}, interface{}) {
	p := tree.root
	if p != nil {
		for p.left != nil {
			p = p.left
		}
	}
	return p.key, p.value
}
func (tree *RedBlackTree) GetMax() (interface{}, interface{}) {
	p := tree.root
	if p != nil {
		for p.right != nil {
			p = p.right
		}
	} else {
		return nil, nil
	}
	return p.key, p.value
}
func (n *Node) printPreOrder(resp *[][]interface{}) {
	if n == nil {
		return
	}

	n.left.printPreOrder(resp)
	*resp = append(*resp, []interface{}{n.key, n.value, n.color})
	n.right.printPreOrder(resp)
}

func (tree *RedBlackTree) GetAll() [][]interface{} {
	res := make([][]interface{}, 0)
	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node == nil {
			return
		}
		inorder(node.left)
		res = append(res, []interface{}{node.key, node.value})
		inorder(node.right)
	}
	inorder(tree.root)
	return res
}

type Entry struct {
	Key   interface{}
	Value interface{}
}
