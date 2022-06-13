package trie

import "strings"

// Trie 字典树，
type Trie struct {
	Root *Node
}

type Node struct {
	End      bool //true 表示这个节点是单词的结尾
	SubNodes map[rune]*Node
}

func (n *Node) addSubNode(k rune, node *Node) {
	if n.SubNodes == nil {
		n.SubNodes = map[rune]*Node{} //初始化SubNodes
	}
	n.SubNodes[k] = node
}

//查看下一个节点有没有对应的字符
func (n *Node) getNextSubNode(k rune) *Node {
	return n.SubNodes[k]
}

func (t *Trie) AddSensitiveWord(word string) {
	tempNode := t.Root
	for i, c := range []rune(word) { //这里需要注意，由于下面（1）用到了len计算word长度，防止中文出错所以[]rune强转
		node := tempNode.getNextSubNode(c)

		if node == nil {
			node = new(Node)
			tempNode.addSubNode(c, node)
		}
		tempNode = node

		if i == len([]rune(word))-1 { //（1）对应这里
			tempNode.End = true
		}
	}
}

//过滤敏感词
//例如：输入 hello,色情不可触碰！
//	   输出 hello,***不可触碰！
func (t *Trie) Filter(words string) string {
	replace := "***"
	tempNode := t.Root

	//拼接字符串的使用
	var build strings.Builder

	for _, c := range words {
		tempNode = tempNode.getNextSubNode(c)
		if tempNode == nil {
			build.WriteRune(c)
			tempNode = t.Root
		} else if tempNode.End {
			build.WriteString(replace)
			tempNode = t.Root
		}
	}
	return build.String()
}

func (t *Trie) ExistSensitiveWord(words string) bool {
	tempNode := t.Root
	for _, c := range words {
		tempNode = tempNode.getNextSubNode(c)
		if tempNode == nil {
			tempNode = t.Root
		} else if tempNode.End {
			return true
		}
	}
	return false
}
