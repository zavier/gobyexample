package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 注意：这里不同于Java，初始化3个时，添加会直接从第四个开始添加
	s = append(s, "x")
	fmt.Println("apd:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d", "e", "f")
	fmt.Println("apd:", s)

	// 切片迭代, 返回的是数据副本
	for index, value := range s {
		fmt.Printf("Index: %d Value: %s\n", index, value)
	}

	// 可以进行复制
	c := make([]string, len(s))
	// 将后面的数据复制到前面，如果不一样大小，则以小的为准
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// 除make外，也可直接声明切片
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 直接声明第100个元素的值，创建长度和容量都为100的切片
	slice2 := []string{99: "x"}
	fmt.Println(slice2)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}