package AhoCorasickDoubleArrayTrie

import (
	"container/list"
	"pangolin/datastructure/rbTree"
)

//An implementation of Aho Corasick algorithm based on Double Array Trie

type AhoCorasickDoubleArrayTrie struct {
	Check  []int           //DAT的check数组
	Base   []int           //base 数组
	Fail   []int           //AC自动机 的fail数组
	Output [][]int         //AC 的output
	Value  []InvertedIndex //固定存储倒排索引
	Len    []int           //每个key的长度
	Size   int             //就是check的数组长度，base数组长度也等于这个

}

func NewAC() *AhoCorasickDoubleArrayTrie {
	return &AhoCorasickDoubleArrayTrie{}
}

//从 map 中导入数据构造AhoCorasickDoubleArrayTrie  map必须是有序的，之后会从其他地方获取就要排序（pdqSort）了，
func (trie *AhoCorasickDoubleArrayTrie) Build(m *rbTree.RedBlackTree) {
	New(trie).Build(m)
}

func (trie *AhoCorasickDoubleArrayTrie) getMatched(pos int, len int, result int, key string, b1 int) int {
	var b = b1
	var p = 0
	for i := pos; i < len; i++ {
		p = b + int(key[i]) + 1
		if b == trie.Check[p] {
			b = trie.Base[p]
		} else {
			return result
		}
	}
	p = b
	n := trie.Base[p]
	if b == trie.Check[p] {
		result = -n - 1
	}
	return result
}

//匹配 text
func (trie *AhoCorasickDoubleArrayTrie) ParseText(text string) *list.List {
	var pos = 1
	var cur = 0
	emits := list.New()
	var runes = []rune(text)
	for i := 0; i < len(runes); i++ {
		cur = trie.getState(cur, runes[i])
		trie.storeEmits(pos, cur, emits)
		pos++
	}
	return emits
}

//相当于增加索引
func (trie *AhoCorasickDoubleArrayTrie) ParseTextWithId(text string, id int) {
	var position = 1
	var currentState = 0

	var runes = []rune(text) //中文特殊处理
	for i := 0; i < len(runes); i++ {
		currentState = trie.getState(currentState, runes[i])
		hitArray := trie.Output[currentState]
		if hitArray != nil {
			for _, hit := range hitArray {
				trie.Value[hit].Value.AddInt(id) //倒排索引的value 使用了roaringBitMap
			}
		}
		position++
	}

}

func (trie *AhoCorasickDoubleArrayTrie) getState(cur int, c rune) int {
	newCur := trie.transWithRoot(cur, c)
	for newCur == -1 {
		cur = trie.Fail[cur]
		newCur = trie.transWithRoot(cur, c)
	}
	return newCur
}

func (trie *AhoCorasickDoubleArrayTrie) storeEmits(pos int, cur int, emits *list.List) {
	hitArr := trie.Output[cur]
	if hitArr != nil {
		for _, hit := range hitArr {
			emits.PushBack(NewHit(pos-trie.Len[hit], pos, trie.Value[hit]))
		}
	}
}

func (trie *AhoCorasickDoubleArrayTrie) transWithRoot(nodePos int, c rune) int {
	b := trie.Base[nodePos]
	p := b + int(c) + 1
	if b != trie.Check[p] {
		if nodePos == 0 {
			return 0
		}
		return -1
	}
	return p
}
