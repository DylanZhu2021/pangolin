package model

import (
	"fmt"
	"pangolin/dao/db"
)

type Doc struct {
	Score int
	Id    int64
	Value *db.Douban
}

// DocSlice 将Interface的方法附加到[]*Doc上，按照**递减**的顺序进行排序。
type DocSlice []*Doc

func (p DocSlice) Len() int           { return len(p) }
func (p DocSlice) Less(i, j int) bool { return p[i].Score > p[j].Score }
func (p DocSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p DocSlice) ShiftTail(a, b int) {
	l := b - a
	if l >= 2 {
		for i := l - 1; i >= 1; i -= 1 {
			if !p.Less(i, i-1) {
				break
			}

			p.Swap(i, i-1)
		}
	}
}

func (p DocSlice) ShiftHead(a, b int) {
	l := b - a
	if l >= 2 {
		for i := 1; i < l; i += 1 {
			if !p.Less(i, i-1) {
				break
			}

			p.Swap(i, i-1)
		}
	}
}

func (p DocSlice) CyclicSwaps(is, js []int) {
	count := len(is)
	tmp := p[is[0]]
	p[is[0]] = p[js[0]]

	for i := 1; i < count; i += 1 {
		p[js[i-1]] = p[is[i]]
		p[is[i]] = p[js[i]]
	}

	p[js[count-1]] = tmp
}

func (d *Doc) Print() {
	fmt.Println()
}
func (d *Doc) PrintString() {
	fmt.Printf("%d - %v", d.Score, d.Value) //打印
	fmt.Println()
}
