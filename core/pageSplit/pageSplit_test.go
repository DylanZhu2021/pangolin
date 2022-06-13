package pageSplit

import (
	"fmt"
	"testing"
)

func TestPageSplit_GetPage(t *testing.T) {
	pageSplit := new(pageSplit)

	var data []int64

	for i := 0; i < 100; i++ {
		data = append(data, int64(i))
	}

	pageSplit.Init(3, 7)
	for i := 1; i <= 3; i++ {
		start, end := pageSplit.GetPage(i)
		fmt.Println(start, end)
	}
}
