/**
 @author: RedCrazyGhost
 @date: 2022/5/31

**/

package dao

import (
	"fmt"
	"pangolin/dao/db"
	"testing"
)

// TestQueryDoubans 注意：--> id []int64 为nil时，查询所有数据
func TestQueryDoubans(t *testing.T) {
	Init()
	id := []int64{1002, 1003} //查询指定key的数据
	id = nil                  //查询所有数据
	doubans, err := db.QueryDoubans(id)
	if err != nil {
		return
	}
	if len(doubans) != 0 {
		for _, value := range doubans {
			fmt.Println(value.Title)
		}
	}
}
