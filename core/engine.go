package core

import (
	"github.com/RoaringBitmap/roaring"
	acdat "pangolin/core/AhoCorasickDoubleArrayTrie"
	"pangolin/core/model"
	"pangolin/core/pageSplit"
	"pangolin/core/sort"
	"pangolin/core/split"
	"pangolin/dao"
	"pangolin/dao/db"
	"pangolin/datastructure/rbTree"
	"pangolin/datastructure/trie"
)

type Engine struct {
	AC      *acdat.AhoCorasickDoubleArrayTrie
	Doubans []*db.Douban
}

func (e *Engine) Init() {
	ac := acdat.NewAC()
	e.AC = ac
	tree := rbTree.NewRedBlackTree()
	dao.Init()
	//doubans, err := db.QueryDoubans([]int64{1001,1002,1003,1004})
	doubans, err := db.QueryDoubans(nil)
	if err != nil {
		panic("数据库出错")
	}
	e.Doubans = doubans
	for _, d := range doubans {
		if d == nil || d.Title == "" {
			continue
		}
		words := split.Seg.CutForSearch(d.Title, true)
		for word := range words {
			if word != "" {
				tree.Add(word, acdat.InvertedIndex{Key: word, Value: roaring.New()})
			}
		}
	}
	ac.Build(tree)

	for _, d := range doubans {
		if d != nil && d.Title != "" {
			ac.ParseTextWithId(d.Title, int(d.Id))
		}
	}

}

// page 从1开始
func (e *Engine) Query(query string, page int, limit int, sensitive string) []*model.Doc {
	rawRes := e.AC.ParseText(query)
	docMap := make(map[int64]*model.Doc)
	bitmap := roaring.New()
	for i := rawRes.Front(); i != nil; i = i.Next() {
		bitmap.Or(i.Value.(*acdat.Hit).Value.Value)
	}
	//fmt.Println(bitmap)
	uint32s := bitmap.ToArray()
	int64s := make([]int64, 0)
	for _, ele := range uint32s {
		int64s = append(int64s, int64(ele))
	}
	doubans, err := db.QueryDoubans(int64s)
	if err != nil {
		panic("数据库出错")
	}
	var tr trie.Trie
	if sensitive != "" {
		tr = trie.Trie{Root: &trie.Node{SubNodes: map[rune]*trie.Node{}}}
		tr.AddSensitiveWord(sensitive)
	}
	for _, d := range doubans {
		if tr == (trie.Trie{}) || !tr.ExistSensitiveWord(d.Title) { //过滤敏感词 修复bug：增加tr对象判空
			docMap[d.Id] = &model.Doc{Id: d.Id, Value: d}
		}
	}
	for i := rawRes.Front(); i != nil; i = i.Next() {
		iterator := i.Value.(*acdat.Hit).Value.Value.Iterator()
		for iterator.HasNext() {
			doc, ok := docMap[int64(iterator.Next())]
			if ok { //有可能因为敏感词被过滤了
				doc.Score += 1 //增加关联度
			}

		}
	}
	res := make([]*model.Doc, 0)
	for _, v := range docMap {
		res = append(res, v)
	}
	sort.Sort(model.DocSlice(res)) //对结果根据关联度排序处理

	newPageSplit := pageSplit.NewPageSplit()
	newPageSplit.Init(limit, len(res))

	start, end := newPageSplit.GetPage(page)

	res0 := make([]*model.Doc, 0)

	for i := start; i < end; i++ {
		res0 = append(res0, res[i])
	}
	return res0
}
