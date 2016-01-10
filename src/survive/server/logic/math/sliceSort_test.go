package math

import (
	"testing"
	"fmt"
	"sort"
)
//定义interface{},并实现sort.Interface接口的三个方法
type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}
func TestSortSlice(t *testing.T){
	a1 := []int{2,4,6,12,353,9,44,244,123}
	a := IntSlice(a1)
	fmt.Println(sort.IsSorted(a)) //false
	if !sort.IsSorted(a) {
		sort.Sort(a)
	}
	fmt.Println(a)
}