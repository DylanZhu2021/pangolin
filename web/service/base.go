package service

import (
	"pangolin/core"
	"pangolin/core/model"
	"pangolin/core/util"
)

type Base struct {
	engine *core.Engine
}

func NewBase(engine *core.Engine) *Base {
	return &Base{engine: engine}
}

// Query 查询
func (b *Base) Query(request *model.SearchRequest) *model.SearchResult {
	result := model.SearchResult{}
	time := util.ExecTime(func() {
		docs := b.engine.Query(request.Query, request.Page, request.Limit, request.Sensitive)
		result.Documents = docs
		result.Total = len(docs)
		result.Limit = request.Limit
		result.Page = request.Page
		result.PageCount = result.Total / result.Limit
		if result.Total%result.Limit != 0 {
			result.PageCount += 1
		}
	})

	result.Time = time
	return &result
}
