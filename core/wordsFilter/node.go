package wordsFilter

import (
	"bytes"
	"strings"
)

type Node struct {
	Child        map[string]*Node
	Placeholders string
}

// NewNode 新建节点
func NewNode(child map[string]*Node, placeholders string) *Node {
	return &Node{
		Child:        child,
		Placeholders: placeholders,
	}
}

//新增敏感词至敏感词map
func (node *Node) add(text string, root map[string]*Node, placeholder string) {
	if text == "" { //若敏感词为空则直接返回
		return
	}
	textr := []rune(text) //int32别名
	end := len(textr) - 1
	for i := 0; i <= end; i++ {
		word := string(textr[i])
		if n, ok := root[word]; ok { //包含当前key
			if i == end { //当前为最后一位
				n.Placeholders = strings.Repeat(placeholder, end+1)
			} else {
				if n.Child != nil {
					root = n.Child
				} else {
					root = make(map[string]*Node)
					n.Child = root
				}
			}
		} else { //如果本身不包含
			placeholders, child := "", make(map[string]*Node)
			if i == end {
				placeholders = strings.Repeat(placeholder, end+1)
			}
			root[word] = NewNode(child, placeholders)
			root = child
		}
	}
}

//移除敏感词从敏感词map中
func (node *Node) remove(text string, root map[string]*Node) {
	textr := []rune(text)
	end := len(textr) - 1
	for i := 0; i <= end; i++ {
		word := string(textr[i])
		if n, ok := root[word]; ok {
			if i == end {
				n.Placeholders = ""
			} else {
				root = n.Child
			}
		} else {
			return
		}
	}
}

// 替换字符串中的敏感词并返回新的字符串。
// 遵循最大限度匹配的原则。
func (node *Node) replace(text string, root map[string]*Node) string {
	if root == nil || text == "" {
		return text
	}
	textr := []rune(text)
	i, s, e, l := 0, 0, 0, len(textr)
	bf := bytes.Buffer{}
	words := make(map[string]*Node)
	var back []*Node
loop:
	for e < l {
		words = root
		i = e
		// 最大匹配原则，首先进行逆向匹配
		for ; i < l; i++ {
			word := string(textr[i])
			if n, ok := words[word]; ok {
				back = append(back, n)
				if n.Child != nil {
					words = n.Child
				} else if n.Placeholders != "" {
					bf.WriteString(string(textr[s:e]))
					bf.WriteString(n.Placeholders)
					i++
					s, e = i, i
					continue loop
				} else {
					break
				}
			} else if n != nil && n.Placeholders != "" {
				bf.WriteString(string(textr[s:e]))
				bf.WriteString(n.Placeholders)
				s, e = i, i
				continue loop
			} else {
				break
			}
		}
		// 后向匹配失败，回溯。
		for ; i > e; i-- {
			bl := len(back)
			if bl == 0 {
				break
			}
			last := back[bl-1]
			back = back[:bl-1]
			if last.Placeholders != "" {
				bf.WriteString(string(textr[s:e]))
				bf.WriteString(last.Placeholders)
				s, e = i, i
				continue loop
			}
		}
		e++
		back = back[:0]
	}
	bf.WriteString(string(textr[s:e]))
	return bf.String()
}

// 查询该字符串是否包含敏感词。
func (node *Node) contains(text string, root map[string]*Node) bool {
	if root == nil || text == "" {
		return false
	}
	textr := []rune(text)
	end := len(textr) - 1
	for i := 0; i <= end; i++ {
		word := string(textr[i])
		if n, ok := root[word]; ok {
			if i == end {
				return n.Placeholders != ""
			} else {
				if len(n.Child) == 0 { //为最后一个
					return true
				}
				root = n.Child
			}
		} else {
			continue
		}
	}
	return false
}
