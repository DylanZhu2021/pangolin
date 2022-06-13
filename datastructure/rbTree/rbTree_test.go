package rbTree

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Add(1, 1)
	tree.Add(2, 10)
	tree.Add(9, 21)
	tree.Add(3, 2)
	tree.Add(5, 22)
	tree.PrintPreOrder()
	fmt.Println(tree.GetMax())
	fmt.Println(tree.GetMin())
	fmt.Println(tree.GetAll())
	for _, v := range tree.GetAll() {
		fmt.Println(v[0])
		fmt.Println(v[1])
	}
}
