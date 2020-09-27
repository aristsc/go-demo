package main //表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
/**
文件夹名与包名没有直接关系，并非需要一致。
同一个文件夹下的文件只能有一个包名，否则编译报错。
*/

import "fmt"

//声明全局变量,允许声明后不被使用
var (
	g int
	h bool
)

func main() {
	fmt.Println("Hello, World!")

	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	var d, e = 123, "hello"
	fmt.Println(b, c, d, e)

	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)

	//空白标识符
	_, numb, strs := number()
	fmt.Println(numb, strs)

	const length, width = 10, 5
	var area int
	area = length * width
	fmt.Printf("面积为: %d\n", area)

	/*  运算符
	&	返回变量存储地址	&a; 将给出变量的实际地址。
	*	指针变量	    	*a; 是一个指针变量*/
	var m = 4
	var ptr *int
	ptr = &m
	fmt.Printf("m 的值为  %d\n", m)
	fmt.Printf("*ptr 为 %d\n", *ptr)
	fmt.Printf("*ptr 为 %d\n", ptr)

	//go语言没有三目运算符
	var o, p = 100, 200
	var ret int
	ret = max(o, p)
	fmt.Printf("最大值是: %d\n", ret)

	r, s := swap("Google", "Runoob")
	fmt.Println(r, s)
}
func number() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//go语言可以返回多个值
func swap(x, y string) (string, string) {
	return y, x
}
