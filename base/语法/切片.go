package main

import "fmt"

/**

切片是对数组对象的抽象
定义切片，并指定长度和容量
make([]T, length, capacity)

*/
func main() {

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	/* 打印原始切片 */
	printSlice(numbers)
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	fmt.Print("\n" +
		"append() 和 copy() 函数\n")
	var numbers5 []int
	printSlice(numbers5)

	/* 允许追加空切片 */
	numbers = append(numbers5, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers5, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers5, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers6 是之前切片的两倍容量*/
	numbers6 := make([]int, len(numbers5), (cap(numbers5))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers6, numbers5)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
