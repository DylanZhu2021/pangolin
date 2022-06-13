package split

import (
	"fmt"
	"github.com/wangbin/jiebago"
)

func init() {
	_ = Seg.LoadDictionary("./dict/dictionary.txt") //加载字典
}

var Seg jiebago.Segmenter

//自定义分词效果打印
func printr(ch <-chan string) {
	for word := range ch {
		fmt.Printf(" %s /", word)
	}
	fmt.Println()
}
