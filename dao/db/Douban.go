/**
 @author: RedCrazyGhost
 @date: 2022/5/31

**/

package db

// Douban 数据类型
type Douban struct {
	Id           int64  `json:"id" gorm:"primaryKey"`
	Title        string `json:"title"`
	Director     string `json:"director"`
	Screenwriter string `json:"screenwriter"`
	Country      string `json:"country"`
	Lang         string `json:"lang"`
	Duration     string `json:"duration"`
	Rank         string `json:"rank"`
	Synopsis     string `json:"synopsis"`
	Url          string `json:"url"`
	Release      string `json:"release"`
	Type         string `json:"type"`
}

// TableName 赋值MySQL表名
func (n *Douban) TableName() string {
	return "douban_movies"
}

// QueryDoubans 查询[]int64内所有id的数据。如果为nil，则查询所有数据。
func QueryDoubans(id []int64) ([]*Douban, error) {
	res := make([]*Douban, 0)
	if err := DB.Find(&res, id).Error; err != nil {
		return nil, err
	}
	return res, nil
}
