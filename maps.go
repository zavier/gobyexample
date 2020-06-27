package main

import "fmt"

func main() {
	// 错误使用方法，需要初始化，如使用make创建
	// var mm map[string]string
	// mm["k"] = "v"

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// 第二个返回值标识对应的键是否存在
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n)
}