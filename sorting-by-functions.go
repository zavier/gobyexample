package main

import (
	"fmt"
	"sort"
)

type byLength []string

// Sort Interface需要实现的3个方法
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// 类型转换
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}