package AhoCorasickDoubleArrayTrie

import "pangolin/datastructure/rbTree"

/**
 * <p>
 * 一个状态有如下几个功能
 * </p>
 * <p/>
 * <ul>
 * <li>success; 成功转移到另一个状态</li>
 * <li>failure; 不可顺着字符串跳转的话，则跳转到一个浅一点的节点</li>
 * <li>emits; 命中一个模式串</li>
 * </ul>
 * <p/>
 * <p>
 * 根节点稍有不同，根节点没有 failure 功能，它的“failure”指的是按照字符串路径转移到下一个状态。其他节点则都有failure状态。
 * </p>
 */

type State struct {
	Depth   int
	Failure *State
	Emits   *rbTree.RedBlackTree //rune
	Success *rbTree.RedBlackTree //rune - State
	Index   int
}

func NewState(depth int) *State {
	s := State{Depth: depth, Emits: rbTree.NewRedBlackTree(), Success: rbTree.NewRedBlackTree()}
	return &s
}

func (s *State) AddEmit(key int) {
	s.Emits.Add(key, struct{}{})
}

func (s *State) GetMax() int {
	if s.Emits == nil || s.Emits.GetTreeSize() == 0 {
		return 0
	}
	k, _ := s.Emits.GetMax()
	return k.(int)
}

func (s *State) isAcceptable() bool {
	return s.Emits != nil && s.Depth > 0
}

func (s *State) nextState(c rune, ignoreRoot bool) *State {
	next := s.Success.Get(c)
	if !ignoreRoot && next == nil && s.Depth == 0 {
		next = s
	}
	if next == nil {
		return nil
	}
	return next.(*State) //注意next == nil 后面就不能加 . 调用了
}

func (s *State) AddState(c rune) *State {
	next := s.nextState(c, true)
	if next == nil {
		next = NewState(s.Depth + 1)
		s.Success.Add(c, next)
	}
	return next
}

func (s *State) SetFailure(f *State, fail []int) {
	s.Failure = f
	fail[s.Index] = f.Index
}
