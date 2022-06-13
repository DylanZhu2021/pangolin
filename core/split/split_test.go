package split

import (
	"fmt"
	"testing"
)

func TestJieba(t *testing.T) {
	fmt.Print("【全模式】：")
	printr(Seg.CutAll("我来到北京清华大学"))

	fmt.Print("【精确模式】：")
	printr(Seg.Cut("我来到北京清华大学", false))

	fmt.Print("【新词识别】：")
	printr(Seg.Cut("他来到了网易杭研大厦", true))

	fmt.Print("【搜索引擎模式】：")
	printr(Seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true)) //我们用的模式
	//【搜索引擎模式】： 小明 / 硕士 / 毕业 / 于 / 中国 / 科学 / 学院 / 科学院 / 中国科学院 / 计算 / 计算所 / ， / 后 / 在 / 日本 / 京都 / 大学 / 日本京都大学 / 深造 /
}
