package main

import "fmt"

func main() {
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i=", i) //1 左移零位，不变
	fmt.Println("j=", j) //j=3(11) 左移三位变为110，即6
	fmt.Println("k=", k) //k=3     左移两位变为1100，即12
	fmt.Println("l=", l) //l=3     左移三位变为11000，即24

	const ( //const只在一个常量组内递增，当有新的const关键字计数会重新开始
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1		在定义常量组时，如果不提供初始值，则表示使用上行的表达式
		h = iota //7,恢复计数
		m        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, m)

}
