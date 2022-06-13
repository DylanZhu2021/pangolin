package AhoCorasickDoubleArrayTrie

import "fmt"

type Hit struct {
	Begin int //在字符串的开始位置
	End   int
	Value InvertedIndex //倒排索引
}

func NewHit(begin int, end int, value InvertedIndex) *Hit {
	hit := Hit{begin, end, value}
	return &hit
}

func (h Hit) PrintString() {
	fmt.Printf("[%d:%d]=%v", h.Begin, h.End, h.Value) //打印
	fmt.Println()
}
