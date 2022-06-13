package trie

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	tr := Trie{Root: &Node{SubNodes: map[rune]*Node{}}}
	tr.AddSensitiveWord("色情")

	fmt.Println(tr.Filter("hello,色情不可触碰！"))

	fmt.Println(tr.ExistSensitiveWord("我讨厌色情"))
	fmt.Println(tr.ExistSensitiveWord("我很纯洁"))
}
