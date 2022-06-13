package pageSplit

import "math"

//分页查询

type pageSplit struct {
	Limit     int //限制显示个数
	PageCount int //总页数
	Total     int //总数据量
}

func NewPageSplit() *pageSplit {
	return &pageSplit{}
}

func (p *pageSplit) Init(limit int, total int) { //对分页初始化
	p.Limit = limit
	p.Total = total

	pageCount := math.Ceil(float64(total) / float64(limit))
	p.PageCount = int(pageCount)
}

func (p *pageSplit) GetPage(currentPage int) (s int, e int) {
	if currentPage > p.PageCount {
		currentPage = p.PageCount
	}

	if currentPage < 1 {
		currentPage = 1
	}

	currentPage -= 1

	//计算位置
	startPosition := currentPage * p.Limit
	endPosition := startPosition + p.Limit

	//若起始位置超过了总量，则默认从0开始
	if startPosition > p.Total {
		return 0, p.Total - 1
	}

	//若结束位置超过了总量，则默认到最后一个的位置
	if endPosition > p.Total {
		endPosition = p.Total
	}

	return startPosition, endPosition
}
