package core

import (
	"testing"
)

func TestInit(t *testing.T) {
	e := new(Engine)
	e.Init()
	query := e.Query("喜羊羊与灰太狼", 1, 10, "牛气冲天")
	for _, q := range query {
		q.PrintString()
	}
	//text := e.AC.ParseText("喜羊羊与灰太狼")

	//fmt.Println(e.AC.Value)
	//fmt.Println("Len:", text.Len())
	//for i := text.Front(); i != nil; i = i.Next() {
	//	i.Value.(*a.Hit).PrintString()
	//}

}
