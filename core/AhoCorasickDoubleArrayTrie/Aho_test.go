package AhoCorasickDoubleArrayTrie

import (
	"fmt"
	"pangolin/datastructure/rbTree"
	"testing"

	"github.com/RoaringBitmap/roaring"
)

func TestBuilder_Build(t *testing.T) {
	ac := NewAC()
	tree := rbTree.NewRedBlackTree()
	tree.Add("我", InvertedIndex{Key: "我", Value: roaring.New()})
	tree.Add("你", InvertedIndex{Key: "你", Value: roaring.New()})
	ac.Build(tree)
	ac.ParseTextWithId("你好", 1)
	ac.ParseTextWithId("我爱我的中国", 2)
	ac.ParseTextWithId("我是谁", 3)
	ac.ParseTextWithId("我爱你", 4)

	text := ac.ParseText("你是真的蠢")

	fmt.Println(ac.Value)
	fmt.Println("Len:", text.Len())
	for i := text.Front(); i != nil; i = i.Next() {
		i.Value.(*Hit).PrintString()
	}
}
