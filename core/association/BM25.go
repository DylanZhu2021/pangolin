package association

//Best Match25

//参考  https://www.elastic.co/guide/cn/elasticsearch/guide/current/pluggable-similarites.html
//https://juejin.cn/post/7012533060398743583
type BM25 struct {
	k1          float32 //这个参数控制着词频结果在词频饱和度中的上升速度。默认值为 1.2 。值越小饱和度变化越快，值越大饱和度变化越慢。
	b           float32 //这个参数控制着字段长归一值所起的作用， 0.0 会禁用归一化， 1.0 会启用完全归一化。默认值为 0.75 。归一化就是使用文档的长度
	freq        float32
	lenOfTarget float32
	aveLen      float32
}

func NewBM25() *BM25 {
	return &BM25{
		k1: 1.2,
		b:  0.75,
	}
}

type BM25Stats struct {
	IdfValue      float32
	Avgdl         float32 //The average document length.
	QueryBoost    float32 //query's inner boost
	TopLevelBoost float32
	Weight        float32
	Cache         []float32
}

func (st *BM25Stats) getValueForNormalization() float32 {
	// we return a TF-IDF like normalization to be nice, but we don't actually normalize ourselves.
	queryWeight := st.IdfValue * st.QueryBoost
	return queryWeight * queryWeight
}

func (st *BM25Stats) normalize(queryNorm float32, topLevelBoost float32) {
	// we don't normalize with queryNorm at all, we just capture the top-level boost
	st.TopLevelBoost = topLevelBoost
	st.Weight = st.IdfValue * st.QueryBoost * topLevelBoost
}

type BM25DocScorer struct {
	Stats       *BM25Stats
	WeightValue float32
	Cache       []float32
}

func (bm *BM25) GetScore(target string, query string) float32 {
	bm.lenOfTarget = float32(len([]rune(target)))

	return ((bm.k1 + 1.0) * bm.freq) / (bm.k1*(1.0-bm.b+bm.b*(bm.lenOfTarget/bm.aveLen)) + bm.freq)
}

//参考 https://blog.csdn.net/ShuiYuanShan/article/details/79871863
func (bm *BM25) getFreq(target string, query string) float32 {
	//Todo  统计词频，这个是放在搜索完成后需要加的功能
	return 0
}
