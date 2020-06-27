package main

import "fmt"

func main() {
	// range 迭代切片
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 遍历Map
	kvs := map[string]string{"a": "apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 遍历字符串
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}