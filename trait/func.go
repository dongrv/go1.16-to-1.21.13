package trait

import "fmt"

func NewFuncs() {
	_ = min(1, 2, 3)
	_ = max(1, 2, 3)

	ints := []int{1, 2, 3, 4, 5, 6}
	clear(ints) // 支持清除：切片、map、类型参数的类型是切片或map
	for i, v := range ints {
		fmt.Printf("ints[%d] = %d\n", i, v) // 应该打印原始长度的0值元素
	}

	maps := map[int]string{0: "0", 1: "1"}
	clear(maps)
	for i, v := range maps {
		fmt.Printf("maps[%d] = %s\n", i, v) // 应该不打印
	}
}
