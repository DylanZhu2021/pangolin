package AhoCorasickDoubleArrayTrie

import (
	"fmt"
	"math"
	"pangolin/datastructure/rbTree"
)
import "pangolin/datastructure/queue"

//用于构造一个AhoCorasickDoubleArrayTrie, 这部分做复杂，尤其是insert
type Builder struct {
	Root      *State
	Used      []bool
	Size      int
	Progress  int
	NextCheck int
	KeySize   int
	trie      *AhoCorasickDoubleArrayTrie
}

func New(trie *AhoCorasickDoubleArrayTrie) *Builder {
	builder := Builder{Root: NewState(0), trie: trie}
	return &builder

}

// m -> string - Index
func (b *Builder) Build(m *rbTree.RedBlackTree) {
	b.trie.Value = make([]InvertedIndex, 0)
	b.trie.Len = make([]int, m.GetTreeSize())
	for i, v := range m.GetAll() {
		b.trie.Value = append(b.trie.Value, v[1].(InvertedIndex))
		b.addKey(v[0].(string), i)
	}
	b.BuildDAT(m.GetTreeSize())
	b.Used = nil
	b.constructFailureStates()
	b.Root = nil
	b.loseWeight()
}

func (b *Builder) addKey(key string, index int) {
	cur := b.Root
	var runes = []rune(key)
	for _, c := range runes {
		cur = cur.AddState(c)
	}
	cur.AddEmit(index)
	b.trie.Len[index] = len(runes)
}

func (b *Builder) BuildDAT(keySize int) {
	b.Progress = 0
	b.KeySize = keySize
	b.Resize(65536 * 32)
	b.trie.Base[0] = 1
	b.NextCheck = 0

	root := b.Root
	siblings := make([]rbTree.Entry, 0)
	b.fetch(root, &siblings)
	if len(siblings) == 0 {
		for i := range b.trie.Check {
			b.trie.Check[i] = -1
		}
	} else {
		b.insert(&siblings)
	}

}

func (b *Builder) fetch(parent *State, siblings *[]rbTree.Entry) int {
	if parent.isAcceptable() {
		fakeNode := NewState(-parent.Depth - 1)
		fakeNode.AddEmit(parent.GetMax())
		*siblings = append(*siblings, rbTree.Entry{Key: rune(0), Value: fakeNode})
	}
	for _, v := range parent.Success.GetAll() {
		*siblings = append(*siblings, rbTree.Entry{Key: v[0].(rune) + 1, Value: v[1].(*State)})
	}
	return len(*siblings)
}

func (b *Builder) Resize(newSize int) int {
	base2 := make([]int, newSize)
	check2 := make([]int, newSize)
	used2 := make([]bool, newSize)
	if b.Size > 0 {
		copy(base2, b.trie.Base)
		copy(check2, b.trie.Check)
		copy(used2, b.Used)
	}
	b.trie.Base = base2
	b.trie.Check = check2
	b.Used = used2
	b.Size = newSize
	return newSize
}
func (b *Builder) loseWeight() {
	nbase := make([]int, b.trie.Size+65535)
	copy(nbase, b.trie.Base)
	b.trie.Base = nbase

	ncheck := make([]int, b.trie.Size+65535)
	copy(ncheck, b.trie.Check)
	b.trie.Check = ncheck

}

func (b *Builder) insert(sibling *[]rbTree.Entry) {
	que := queue.New()
	que.Push(rbTree.Entry{Key: nil, Value: sibling})
	for !que.Empty() {
		cur := que.Pop().(rbTree.Entry)
		siblings := *(cur.Value.(*[]rbTree.Entry))
		var begin = 0
		var pos = int(math.Max(float64((siblings[0].Key.(rune))+1), float64(b.NextCheck)) - 1)
		var first = 0
		var nonzero_num = 0
		if b.Size <= pos {
			b.Resize(pos + 1)
		}
	outer:
		for {
			pos++

			if b.Size <= pos {
				b.Resize(pos + 1)
			}
			if b.trie.Check[pos] != 0 {
				nonzero_num++
				continue
			} else if first == 0 {
				b.NextCheck = pos
				first = 1
			}

			begin = pos - int(siblings[0].Key.(rune)) // 当前位置离第一个兄弟节点的距离
			if b.Size <= (begin + int(siblings[len(siblings)-1].Key.(rune))) {
				// progress can be zero // 防止progress产生除零错误
				toSize := math.Max(1.05, float64(1.0*b.KeySize/(b.Progress+1))) * float64(b.Size)
				maxSize := 0x7fffffff * 0.95
				if float64(b.Size) >= maxSize {
					fmt.Println("Error:Double array trie is too big. -line150")
				} else {
					b.Resize(int(math.Min(toSize, maxSize)))
				}
			}

			if b.Used[begin] {
				continue
			}
			for i, sib := range siblings {
				if i == 0 {
					continue
				}
				if b.trie.Check[begin+int(sib.Key.(rune))] != 0 {
					continue outer
				}
			}
			break
		}
		if float64(1.0*nonzero_num/(pos-b.NextCheck+1)) >= 0.95 {
			b.NextCheck = pos // 从位置 next_check_pos 开始到 pos 间，如果已占用的空间在95%以上，下次插入节点时，直接从 pos 位置处开始查找
		}
		b.Used[begin] = true

		b.trie.Size = int(math.Max(float64(b.trie.Size), float64(begin+int(siblings[len(siblings)-1].Key.(rune))+1)))

		for _, sib := range siblings {
			b.trie.Check[begin+int(sib.Key.(rune))] = begin
		}
		for _, sib := range siblings {
			new_siblings := make([]rbTree.Entry, 0) //Entry<Integer,State>
			if b.fetch(sib.Value.(*State), &new_siblings) == 0 {
				b.trie.Base[begin+int(sib.Key.(rune))] = -sib.Value.(*State).GetMax() - 1
				b.Progress++
			} else {
				tmp := rbTree.Entry{Key: rune(begin) + sib.Key.(rune), Value: &new_siblings}
				que.Push(tmp)
			}
			sib.Value.(*State).Index = begin + int(sib.Key.(rune))
		}

		if cur.Key != nil {
			b.trie.Base[cur.Key.(rune)] = begin
		}

	}
}

func (b *Builder) constructFailureStates() {
	b.trie.Fail = make([]int, b.trie.Size+1)
	b.trie.Output = make([][]int, b.trie.Size+1)
	qu := queue.New() //<State>
	// 第一步，将深度为1的节点的failure设为根节点
	for _, v := range b.Root.Success.GetAll() {
		v[1].(*State).SetFailure(b.Root, b.trie.Fail)
		qu.Push(v[1])
		b.constructOutput(v[1].(*State))
	}
	for !qu.Empty() {
		cur := qu.Pop().(*State)

		for _, v := range cur.Success.GetAll() {
			nextState := cur.nextState(v[0].(rune), false)
			qu.Push(nextState)

			failure := cur.Failure
			for failure.nextState(v[0].(rune), false) == nil {
				failure = failure.Failure
			}
			newFail := failure.nextState(v[0].(rune), false)
			nextState.SetFailure(newFail, b.trie.Fail)
			for _, v := range newFail.Emits.GetAll() {
				nextState.AddEmit(v[0].(int))
			}
			b.constructOutput(nextState)
		}
	}
}

func (b *Builder) constructOutput(target *State) {
	emits := target.Emits
	if emits == nil || emits.GetTreeSize() == 0 {
		return
	}
	output := make([]int, emits.GetTreeSize())
	var i = 0
	for _, v := range emits.GetAll() {
		output[i] = v[0].(int)
		i++
	}
	b.trie.Output[target.Index] = output
}
