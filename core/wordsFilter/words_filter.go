package wordsFilter

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
)

var DefaultPlaceholder = "*"
var DefaultStripSpace = true

type WordsFilter struct {
	Placeholder string
	StripSpace  bool
	node        *Node
	mutex       sync.RWMutex
}

// New 新建一个敏感词过滤器
func New() *WordsFilter {
	return &WordsFilter{
		Placeholder: DefaultPlaceholder,
		StripSpace:  DefaultStripSpace,
		node:        NewNode(make(map[string]*Node), ""),
	}
}

// Generate 将敏感文本列表转换为敏感词树节点
func (wf *WordsFilter) Generate(texts []string) map[string]*Node {
	root := make(map[string]*Node)
	for _, text := range texts {
		wf.Add(text, root)
	}
	return root
}

// GenerateWithFile 将文件中的敏感文本转换成敏感词树节点。
// 文件内容格式，请包住每一个敏感词。
func (wf *WordsFilter) GenerateWithFile(path string) (map[string]*Node, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	buf := bufio.NewReader(fd)
	var texts []string
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		text := strings.TrimSpace(string(line)) //去除空格
		if text == "" {
			continue
		}
		texts = append(texts, text)
	}

	root := wf.Generate(texts)
	return root, nil
}

// Add 将敏感词添加到指定的敏感词中
func (wf *WordsFilter) Add(text string, root map[string]*Node) {
	if wf.StripSpace {
		text = stripSpace(text)
	}
	wf.mutex.Lock()
	defer wf.mutex.Unlock()
	wf.node.add(text, root, wf.Placeholder)
}

// Replace 替换敏感词为新的词
func (wf *WordsFilter) Replace(text string, root map[string]*Node) string {
	if wf.StripSpace {
		text = stripSpace(text)
	}
	wf.mutex.RLock()
	defer wf.mutex.RUnlock()
	return wf.node.replace(text, root)
}

// Contains 查看是否包含敏感词
func (wf *WordsFilter) Contains(text string, root map[string]*Node) bool {
	if wf.StripSpace {
		text = stripSpace(text)
	}
	wf.mutex.RLock()
	defer wf.mutex.RUnlock()
	return wf.node.contains(text, root)
}

// Remove 去除敏感词
func (wf *WordsFilter) Remove(text string, root map[string]*Node) {
	if wf.StripSpace {
		text = stripSpace(text)
	}
	wf.mutex.Lock()
	defer wf.mutex.Unlock()
	wf.node.remove(text, root)
}

//剥离空间
func stripSpace(str string) string {
	fields := strings.Fields(str)
	var bf bytes.Buffer
	for _, field := range fields {
		bf.WriteString(field)
	}
	return bf.String()
}
